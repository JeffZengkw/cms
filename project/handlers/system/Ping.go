package system

import (
	"fmt"
	"gotest/project/utils"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	fmt.Println(utils.DB())
	ctx.JSON(200, gin.H{
		"msg": "pong",
	})
}
