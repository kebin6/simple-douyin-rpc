package config

import (
	DYConfig "github.com/kebin6/simple-douyin-rpc/open/config"
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	RedisConf    config.RedisConf
	DouYinConf   DYConfig.Config
}
