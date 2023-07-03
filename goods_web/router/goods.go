package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mingshop_api/goods_web/api/goods"
	"mingshop_api/goods_web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("/goods").Use(middlewares.Trace())
	zap.S().Info("配置商品路由")
	{
		GoodsRouter.GET("", goods.List)                                                                 //获取商品列表
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)               //添加商品
		GoodsRouter.GET("/:id", goods.Detail)                                                           //获取商品详情
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete)      //删除商品
		GoodsRouter.GET("/:id/stocks", goods.Stocks)                                                    //获取商品库存
		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)         //更新商品
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus) //更新商品部分状态
	}
}
