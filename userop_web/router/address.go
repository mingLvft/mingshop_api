package router

import (
	"github.com/gin-gonic/gin"
	"mingshop_api/userop_web/api/address"
	"mingshop_api/userop_web/middlewares"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middlewares.JWTAuth(), address.List)          // 地址列表页
		AddressRouter.DELETE("/:id", middlewares.JWTAuth(), address.Delete) // 删除地址
		AddressRouter.POST("", middlewares.JWTAuth(), address.New)          //新建地址
		AddressRouter.PUT("/:id", middlewares.JWTAuth(), address.Update)    //修改地址信息
	}
}
