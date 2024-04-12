package config

import "github.com/kebin6/simple-douyin-rpc/open/cache"

// Config for 抖音开放平台
type Config struct {
	ClientKey    string      `json:",env=DY_CLIENT_KEY"`
	ClientSecret string      `json:",env=DY_CLIENT_SECRET"`
	RedirectURL  string      `json:",optional"`
	Scopes       string      `json:",optional"`
	Cache        cache.Cache `json:",optional"`
}
