package global

import (
	ut "github.com/go-playground/universal-translator"
	"mingshop_api/goods_web/config"
	"mingshop_api/goods_web/proto"
)

var (
	Trans          ut.Translator
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	NacosConfig    *config.NacosConfig  = &config.NacosConfig{}
	GoodsSrvClient proto.GoodsClient
)
