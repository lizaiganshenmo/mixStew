// Code generated by hertz generator.

package main

import (
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/rpc"
)

func Init() {
	// rpc init
	rpc.Init()
	// log init
	logInit()
}

func logInit() {
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	hlog.SetOutput(f)
}

func main() {
	Init()

	h := server.Default()

	register(h)
	h.Spin()
}
