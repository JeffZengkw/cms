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
	//设置应用的模式
	utils.SetMode() //发布模式，后台会少打印很多信息
	//初始化日志
	utils.SetLogger()
	//初始化数据库连接
	utils.InitDB()

	//初始化路由引擎
	r := handlers.InitEngine()

	//使用 Logger， 输出应用日志
	utils.Logger().Info("service is listening on ", "addr", viper.GetString("app.addr"))
	r.Run(viper.GetString("app.addr")) //使用配置文件中的数据
}
