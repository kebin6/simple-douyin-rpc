package dyApi

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/kebin6/simple-douyin-rpc/open"
	"github.com/kebin6/simple-douyin-rpc/open/credential"
	"github.com/kebin6/simple-douyin-rpc/open/util"
	"strconv"
	"time"

	"github.com/kebin6/simple-douyin-rpc/internal/svc"
	"github.com/kebin6/simple-douyin-rpc/types/douyin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSignatureLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSignatureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSignatureLogic {
	return &GetSignatureLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSignatureLogic) GetSignature(in *douyin.SignatureReq) (*douyin.SignatureResp, error) {
	dyCache := l.svcCtx.NewRedisCache(l.ctx, l.svcCtx.Redis)

	dyConfig := l.svcCtx.Config.DouYinConf
	dyConfig.Cache = dyCache
	dyApi := open.NewOpenAPI(&dyConfig)
	clientToken, err := dyApi.GetClientToken()
	if err != nil {
		return nil, err
	}

	douYinJSTicketHandler := credential.NewDefaultJsTicket(
		l.svcCtx.Config.DouYinConf.ClientKey, "douyin", dyCache)
	jsTicket, err := douYinJSTicketHandler.GetTicket(clientToken)
	if err != nil {
		return nil, err
	}

	sign, nonceStr, timestamp := l.Sign(jsTicket, in.Url)
	return &douyin.SignatureResp{
		Signature: sign,
		NonceStr:  nonceStr,
		Timestamp: timestamp,
		ClientKey: l.svcCtx.Config.DouYinConf.ClientKey,
	}, nil
}

func (l *GetSignatureLogic) Sign(jsapiTicket string, url *string) (sign string, nonceStr string, timestamp int64) {
	nonceStr = util.GenRandomString(16)
	timestamp = time.Now().Unix()
	// 对所有待签名参数按照字段名的 ASCII 码从小到大排序（字典序）后，使用 URL 键值对的格式（即 key1=value1&key2=value2…）拼接成字符串 string1
	string1 := "jsapi_ticket=" + jsapiTicket + "&noncestr=" + nonceStr + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&url=" + *url
	// 对 string1 进行 MD5 签名，得到 signature
	hash := md5.Sum([]byte(string1))
	return hex.EncodeToString(hash[:]), nonceStr, timestamp
}
