package global

import (
	ut "github.com/go-playground/universal-translator"
	"mingshop_api/userop_web/config"
	"mingshop_api/userop_web/proto"
)

var (
	Trans          ut.Translator
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	NacosConfig    *config.NacosConfig  = &config.NacosConfig{}
	GoodsSrvClient proto.GoodsClient

	MessageClient proto.MessageClient
	AddressClient proto.AddressClient
	UserFavClient proto.UserFavClient
)
