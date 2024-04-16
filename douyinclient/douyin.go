// Code generated by goctl. DO NOT EDIT.
// Source: douyin.proto

package douyinclient

import (
	"context"

	"github.com/kebin6/simple-douyin-rpc/types/douyin"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccessTokenReq   = douyin.AccessTokenReq
	AccessTokenResp  = douyin.AccessTokenResp
	BaseIDInt32Resp  = douyin.BaseIDInt32Resp
	BaseIDInt64Resp  = douyin.BaseIDInt64Resp
	BaseIDResp       = douyin.BaseIDResp
	BaseIDUint32Resp = douyin.BaseIDUint32Resp
	BaseResp         = douyin.BaseResp
	BaseUUIDResp     = douyin.BaseUUIDResp
	Empty            = douyin.Empty
	IDInt32Req       = douyin.IDInt32Req
	IDInt64Req       = douyin.IDInt64Req
	IDReq            = douyin.IDReq
	IDUint32Req      = douyin.IDUint32Req
	IDsInt32Req      = douyin.IDsInt32Req
	IDsInt64Req      = douyin.IDsInt64Req
	IDsReq           = douyin.IDsReq
	IDsUint32Req     = douyin.IDsUint32Req
	PageInfoReq      = douyin.PageInfoReq
	SignatureReq     = douyin.SignatureReq
	SignatureResp    = douyin.SignatureResp
	UUIDReq          = douyin.UUIDReq
	UUIDsReq         = douyin.UUIDsReq

	Douyin interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		GetSignature(ctx context.Context, in *SignatureReq, opts ...grpc.CallOption) (*SignatureResp, error)
		GetAccessToken(ctx context.Context, in *AccessTokenReq, opts ...grpc.CallOption) (*AccessTokenResp, error)
	}

	defaultDouyin struct {
		cli zrpc.Client
	}
)

func NewDouyin(cli zrpc.Client) Douyin {
	return &defaultDouyin{
		cli: cli,
	}
}

func (m *defaultDouyin) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := douyin.NewDouyinClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

func (m *defaultDouyin) GetSignature(ctx context.Context, in *SignatureReq, opts ...grpc.CallOption) (*SignatureResp, error) {
	client := douyin.NewDouyinClient(m.cli.Conn())
	return client.GetSignature(ctx, in, opts...)
}

func (m *defaultDouyin) GetAccessToken(ctx context.Context, in *AccessTokenReq, opts ...grpc.CallOption) (*AccessTokenResp, error) {
	client := douyin.NewDouyinClient(m.cli.Conn())
	return client.GetAccessToken(ctx, in, opts...)
}
