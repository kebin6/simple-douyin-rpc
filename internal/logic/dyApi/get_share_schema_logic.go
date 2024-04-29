package dyApi

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/kebin6/simple-douyin-rpc/internal/svc"
	"github.com/kebin6/simple-douyin-rpc/open"
	"github.com/kebin6/simple-douyin-rpc/open/credential"
	"github.com/kebin6/simple-douyin-rpc/open/share"
	"github.com/kebin6/simple-douyin-rpc/open/util"
	"github.com/kebin6/simple-douyin-rpc/types/douyin"
	"github.com/skip2/go-qrcode"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	schemaUrl = "snssdk1128://openplatform/share?share_type=%s&client_key=%s&nonce_str=%s&timestamp=%s&signature=%s&state=%s&image_path=%s&image_list_path=%s&video_path=%s&hashtag_list=%s&micro_app_info=%s&share_to_publish=%d&title=%s&poi_id=%s"
)

type GetShareSchemaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShareSchemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareSchemaLogic {
	return &GetShareSchemaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShareSchemaLogic) GetShareSchema(in *douyin.GetShareSchemaReq) (*douyin.GetShareSchemaResp, error) {
	dyCache := l.svcCtx.NewRedisCache(l.ctx, l.svcCtx.Redis)

	dyConfig := l.svcCtx.Config.DouYinConf
	dyConfig.Cache = dyCache
	dyApi := open.NewOpenAPI(&dyConfig)
	clientToken, err := dyApi.GetClientToken()
	if err != nil {
		return nil, err
	}
	// 获取 share_id，以便后续可以获取视频发布情况
	shareId, err := share.NewDefaultShareId(clientToken, dyCache).GetShareId()
	if err != nil {
		return nil, err
	}
	// 获取签名
	openTicket, err := l.GetOpenTicket()
	if err != nil {
		return nil, err
	}
	signature, noceStr, timestamp := l.Sign(*openTicket)

	shareType := "h5"
	clientKey := dyConfig.ClientKey
	state := shareId
	imagePath := ""
	if in.ImagePath != nil {
		imagePath = *in.ImagePath
	}
	imageListPath := ""
	if in.ImageListPath != nil && len(in.ImageListPath) > 1 {
		imageListPathJsonBytes, err := json.Marshal(in.ImageListPath)
		if err != nil {
			return nil, err
		}
		imageListPath = string(imageListPathJsonBytes)

		//params, err := url.ParseQuery(imageListPath)
		//if err != nil {
		//	return nil, err
		//}
		//imageListPath = params.Encode()
	}
	videoPath := ""
	if in.VideoPath != nil {
		videoPath = *in.VideoPath
	}
	hashtagList := ""
	if in.HashtagList != nil && len(in.HashtagList) > 1 {
		hashtagListJsonBytes, err := json.Marshal(in.HashtagList)
		if err != nil {
			return nil, err
		}
		hashtagList = string(hashtagListJsonBytes)

		//params, err := url.ParseQuery(hashtagList)
		//if err != nil {
		//	return nil, err
		//}
		//hashtagList = params.Encode()
	}
	microAppInfo := ""
	if in.MicroAppInfo != nil && *in.MicroAppInfo != "" {
		microAppInfo = *in.MicroAppInfo
	}
	shareToPublish := 0
	if in.ShareToPublish {
		shareToPublish = 1
	}
	title := in.Title
	poiId := ""
	if in.PoiId != nil && *in.PoiId != "" {
		poiId = *in.PoiId
	}
	schema := fmt.Sprintf(schemaUrl, shareType, clientKey, noceStr, timestamp, signature, state, imagePath, imageListPath, videoPath, hashtagList, microAppInfo, shareToPublish, title, poiId)

	// 生成二维码
	//var qrcodeImage []byte
	//qrcodeImage, err = qrcode.Encode(schema, qrcode.Medium, 256)
	//if err != nil {
	//	return nil, err
	//}
	//schema = "data:image/png;base64," + base64.StdEncoding.EncodeToString(qrcodeImage)

	// 生成二维码图片
	err = qrcode.WriteFile(schema, qrcode.Medium, 256, "qr.png")
	if err != nil {
		return nil, err
	}
	return &douyin.GetShareSchemaResp{
		Schema:  schema,
		ShareId: shareId,
	}, nil
}

func (l *GetShareSchemaLogic) GetOpenTicket() (*string, error) {
	dyCache := l.svcCtx.NewRedisCache(l.ctx, l.svcCtx.Redis)

	dyConfig := l.svcCtx.Config.DouYinConf
	dyConfig.Cache = dyCache
	dyApi := open.NewOpenAPI(&dyConfig)
	clientToken, err := dyApi.GetClientToken()
	if err != nil {
		return nil, err
	}

	douYinOpenTicketHandler := credential.NewDefaultOpenTicket(
		l.svcCtx.Config.DouYinConf.ClientKey, "douyin", dyCache)
	openTicket, err := douYinOpenTicketHandler.GetTicket(clientToken)
	if err != nil {
		return nil, err
	}
	return &openTicket, nil
}

func (l *GetShareSchemaLogic) Sign(openTicket string) (sign string, nonceStr string, timestampStr string) {
	nonceStr = util.GenRandomString(16)
	timestamp := time.Now().Unix()
	timestampStr = strconv.FormatInt(timestamp, 10)
	// 对所有待签名参数按照字段名的 ASCII 码从小到大排序（字典序）后，使用 URL 键值对的格式（即 key1=value1&key2=value2…）拼接成字符串 string1
	string1 := "nonce_str=" + nonceStr + "&ticket=" + openTicket + "&timestamp=" + timestampStr
	// 对 string1 进行 MD5 签名，得到 signature
	hash := md5.Sum([]byte(string1))
	return hex.EncodeToString(hash[:]), nonceStr, timestampStr
}
