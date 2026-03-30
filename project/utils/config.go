package utils

import (
	"log"

	"github.com/spf13/viper"
)

// 默认配置
func defaultConfig() {
	viper.SetDefault("app.addr", ":8080")
}

//编辑配置工具

func ParseConfig() {
	//默认配置
	defaultConfig() //调用函数完成
	//配置解析参数
	viper.AddConfigPath("./project") //从哪些目录搜索config文件
	viper.SetConfigName("configs")   //文件名和后缀分开，这里填文件名，下一个 type填格式
	viper.SetConfigType("yaml")      //配置类型（格式）

	//执行解析，调用viper的ReadInConfig()方法
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)

	}

}
