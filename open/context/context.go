package context

import (
	"github.com/kebin6/simple-douyin-rpc/open/config"
	"github.com/kebin6/simple-douyin-rpc/open/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
