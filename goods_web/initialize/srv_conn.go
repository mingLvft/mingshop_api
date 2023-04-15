package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mingshop_api/goods_web/global"
	"mingshop_api/goods_web/proto"
)

func InitSrvConn() {
	//从注册中心获取获取服务的信息，并对grpc与consul建立连接获取服务，然后使用gprc的负载均衡（服务是动态的）
	//grpc内部提供了接口实现负载均衡和服务发现（只提供了接口，可以指定注册中心/DNS，去获取服务，需要grpc-consul-resolver第三方库完成获取操作）
	//grpc-consul-resolver实现consul获取服务到grpc。grpc实现了连接的负载均衡
	consulInfo := global.ServerConfig.ConsulInfo
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【商品服务失败】")
	}

	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)
}
