package main

import (
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	config "github.com/lizaiganshenmo/mixStew/cmd/follow/configs"
	"github.com/lizaiganshenmo/mixStew/cmd/follow/dal"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/follow/followservice"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	"github.com/lizaiganshenmo/mixStew/library/middleware"
)

func Init() {
	// 配置信息加载
	config.Init()
	// dal 初始化
	dal.Init()
	// log init
	logInit()

}

func logInit() {
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// defer f.Close()
	klog.SetOutput(f)

}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAdd})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.FollowServiceListenADD)

	svr := followservice.NewServer(
		new(FollowServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.FollowServiceName,
		}),
		server.WithMiddleware(middleware.CommonMiddleware), // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: constants.MaxConnections, MaxQPS: constants.MaxQPS}),
	)
	if err := svr.Run(); err != nil {
		panic(err)
	}
}
