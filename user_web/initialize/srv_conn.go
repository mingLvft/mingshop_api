package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mingshop_api/user_web/global"
	"mingshop_api/user_web/proto"
)

func InitSrvConn() {
	//从注册中心获取获取服务的信息，并对grpc与consul建立连接获取服务，然后使用gprc的负载均衡（服务是动态的）
	//grpc内部提供了接口实现负载均衡和服务发现（只提供了接口，可以指定注册中心/DNS，去获取服务，需要grpc-consul-resolver第三方库完成获取操作）
	//grpc-consul-resolver实现consul获取服务到grpc。grpc实现了连接的负载均衡
	consulInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func InitSrvConn2() {
	//从注册中心获取获取服务的信息(服务是固定的)
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		return
	}

	//拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 【用户服务失败】",
			"msg", err.Error(),
		)
	}
	//1. 后续的用户服务下线了 2. 改端口了 3.改ip了
	//事先创立好了连接，这样后续就不用进行再次的tcp三次握手
	//一个连接多个groutine共用，会不会出现性能问题，这里需要做一个连接池。负载均衡也能解决这个问题
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
