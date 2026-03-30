package handlers

import (
	"gotest/project/handlers/system"

	"github.com/gin-gonic/gin"
)

//初始化路由引擎

func InitEngine() *gin.Engine {

	//1、初始化路由引擎
	r := gin.Default()
	//下面这一行路由在 业务逻辑中去实现
	//r.GET("/ping", system.Ping)
	//init中只需要注册不同模块的路由

	//2、注册不同模块的路由
	system.Router(r)

	//3、返回注册好的路由数据
	return r
}
