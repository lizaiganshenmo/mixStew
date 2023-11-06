package main

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	config "github.com/lizaiganshenmo/mixStew/cmd/user/configs"
	"github.com/lizaiganshenmo/mixStew/cmd/user/dal"
	"github.com/lizaiganshenmo/mixStew/cmd/user/resource"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user/userservice"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	"github.com/lizaiganshenmo/mixStew/library/middleware"
)

func Init() {
	// 配置信息加载-应当在首位顺序
	config.Init()
	// dal 初始化
	dal.Init()
	// resource 初始化 全局变量初始化 包括 log
	resource.Init()

}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAdd})

	if err != nil {
		panic(err)
	}

	// add tracer
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.UserServiceName),
		provider.WithExportEndpoint(constants.JaegerColAdd),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)
	defer p.Shutdown(context.Background())

	addr, err := net.ResolveTCPAddr("tcp", constants.UserServiceListenADD)

	svr := userservice.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.UserServiceName,
		}),
		server.WithMiddleware(middleware.CommonMiddleware), // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: constants.MaxConnections, MaxQPS: constants.MaxQPS}),
		server.WithSuite(tracing.NewServerSuite()), // tracer
	)
	if err := svr.Run(); err != nil {
		panic(err)
	}
}
