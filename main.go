package main

import (
	"web-server/config"
	"web-server/router"
)

func main() {
	//初始化数据库配置
	config.InitConfig()
	//设置路由
	r := router.NewRouter()
	r.Run(":4000")

}
