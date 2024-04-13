package main

import (
	"fmt"
	"g-shopping/user-web/initialize"
	"go.uber.org/zap"
)

var (
	port = 5621
)

func main() {
	initialize.InitialLogger()
	r := initialize.Routers()

	zap.S().Infof("启动服务器， 端口：%d", port)

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Error("启动服务器失败 ", err.Error())
	}
}
