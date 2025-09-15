package main

import (
	"github.com/yhcui/web3study/task4/global"
	"github.com/yhcui/web3study/task4/initialize"
)

func main() {
	global.SDB = initialize.GROM()
	router := initialize.Routes()
	router.Run(":8080")

}
