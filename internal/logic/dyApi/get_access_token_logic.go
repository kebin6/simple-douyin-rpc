package dyApi

import (
	"context"
	"github.com/kebin6/simple-douyin-rpc/open"

	"github.com/kebin6/simple-douyin-rpc/internal/svc"
	"github.com/kebin6/simple-douyin-rpc/types/douyin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccessTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccessTokenLogic {
	return &GetAccessTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccessTokenLogic) GetAccessToken(in *douyin.AccessTokenReq) (*douyin.AccessTokenResp, error) {
	dyCache := l.svcCtx.NewRedisCache(l.ctx, l.svcCtx.Redis)

	dyConfig := l.svcCtx.Config.DouYinConf
	dyConfig.Cache = dyCache
	accessTokenResp, err := open.NewOpenAPI(&dyConfig).GetOauth().GetUserAccessToken(in.GetCode())
	if err != nil {
		return nil, err
	}
	return &douyin.AccessTokenResp{
		AccessToken: accessTokenResp.AccessToken,
	}, nil
}
