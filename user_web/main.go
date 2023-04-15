package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mingshop_api/user_web/global"
	"mingshop_api/user_web/initialize"
	"mingshop_api/user_web/utils"
	"mingshop_api/user_web/utils/register/consul"
	mingshopvalidator "mingshop_api/user_web/validator"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化logger
	initialize.InitLogger()

	// 初始化routers
	Router := initialize.Routers()

	//初始化配置文件
	initialize.InitConfig()
	zap.S().Infof("nacos配置信息：%v", global.ServerConfig)

	//初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	//初始化srv的连接
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	//如果是本地开发环境端口号固定，线上环境自动获取端口号
	debug := viper.GetBool("MINGSHOP_DEBUG")
	if !debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	//注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", mingshopvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}必须是有效的手机号", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	// S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
	// 日志分级别 debug,info, warn, error, panic, fatal
	//服务注册
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败：", err.Error())
	}
	zap.S().Debugf("启动服务器，监听端口：%d", global.ServerConfig.Port)
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()
	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = register_client.DeRegister(serviceId); err != nil {
		zap.S().Info("注销服务失败")
	} else {
		zap.S().Info("注销服务成功")
	}
}
