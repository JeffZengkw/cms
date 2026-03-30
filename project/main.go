package main

import (
	"gotest/project/handlers"
	"gotest/project/utils"

	"github.com/spf13/viper"
)

//现在目录的入口文件

func main() {
	//解析配置
	utils.ParseConfig()

	//初始化路由引擎
	r := handlers.InitEngine()
	r.Run(viper.GetString("app.addr")) //使用配置文件中的数据
}
