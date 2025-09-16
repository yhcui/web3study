package main

import (
	"github.com/yhcui/web3study/task4/global"
	"github.com/yhcui/web3study/task4/initialize"
)

func main() {
	global.Logger = initialize.InitLogger()
	global.SDB = initialize.GROM()
	router := initialize.Routers()

	global.Logger.Info("服务准备启动")
	router.Run(":8080")

}
