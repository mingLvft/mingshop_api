package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"mingshop_api/order_web/global"
)

func Trace() gin.HandlerFunc {
	return func(context *gin.Context) {
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: fmt.Sprintf("%s:%d", global.ServerConfig.JaegerInfo.Host, global.ServerConfig.JaegerInfo.Port),
			},
			ServiceName: global.ServerConfig.JaegerInfo.Name,
		}

		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
		if err != nil {
			panic(err)
		}
		opentracing.SetGlobalTracer(tracer)
		defer closer.Close()

		startSpan := tracer.StartSpan(context.Request.URL.Path)
		defer startSpan.Finish()

		// 在gin路由中进行了创建trace和span，但是只针对了每个api请求，无法对其中的一个或者多个grpc调用进行追踪
		context.Set("tracer", tracer)
		context.Set("parentSpan", startSpan)
		context.Next()
	}
}
