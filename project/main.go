package main

import (
	"gotest/project/handlers/system"

	"github.com/gin-gonic/gin"
)

//现在目录的入口文件

func main() {
	r := gin.Default()

	r.GET("/ping", system.Ping)
	r.Run()
}
