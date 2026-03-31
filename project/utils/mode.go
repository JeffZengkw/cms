package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 设置应用模式
func SetMode() {
	switch strings.ToLower(viper.GetString("app.mode")) { //strings.ToLower 是为了避免 大小写问题，统一转换为 小写
	case "release":
		gin.SetMode(gin.ReleaseMode) //gin框架中的 模式转发方法
	case "test":
		gin.SetMode(gin.TestMode)
	case "debug":
		fallthrough //fallthrough 只会强制跳到【紧邻的下一个 case】，不是最后一个，也不是跳过所有！
	default:
		gin.SetMode(gin.DebugMode)
	}

}
