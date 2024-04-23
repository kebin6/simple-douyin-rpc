package dyApi

import (
	"context"
	"fmt"
	"github.com/kebin6/simple-douyin-rpc/open"
	"github.com/kebin6/simple-douyin-rpc/open/share"

	"github.com/kebin6/simple-douyin-rpc/internal/svc"
	"github.com/kebin6/simple-douyin-rpc/types/douyin"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	schemaUrl = "snssdk1128://openplatform/share?share_type=%s&client_key=%s&nonce_str=%s&timestamp=%d&signature=%s&state=%s&image_path=%s&image_list_path=%s&video_path=%s&hashtag_list=%s&micro_app_info=%s&share_to_publish=%s&title=%s&poi_id=%s"
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
	signatureLogic := NewGetSignatureLogic(l.ctx, l.svcCtx)
	signatureResp, err := signatureLogic.GetSignature(&douyin.SignatureReq{})
	if err != nil {
		return nil, err
	}

	shareType := "h5"
	clientKey := dyConfig.ClientKey
	noceStr := signatureResp.NonceStr
	timestamp := signatureResp.Timestamp
	signature := signatureResp.Signature
	state := shareId
	imagePath := in.ImagePath
	imageListPath := in.ImageListPath
	videoPath := in.VideoPath
	hashtagList := in.HashtagList
	microAppInfo := in.MicroAppInfo
	shareToPublish := in.ShareToPublish
	title := in.Title
	poiId := in.PoiId

	schema := fmt.Sprintf(schemaUrl, shareType, clientKey, noceStr, timestamp, signature, state, imagePath, imageListPath, videoPath, hashtagList, microAppInfo, shareToPublish, title, poiId)
	return &douyin.GetShareSchemaResp{
		Schema: schema,
	}, nil
}
