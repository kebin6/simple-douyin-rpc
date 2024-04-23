package dyApi

import (
	"context"

	"github.com/kebin6/simple-douyin-rpc/internal/svc"
	"github.com/kebin6/simple-douyin-rpc/types/douyin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *douyin.UserInfoReq) (*douyin.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &douyin.UserInfoResp{}, nil
}
