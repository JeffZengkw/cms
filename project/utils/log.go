package utils

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//设置日志的函数

func SetLogger() {
	//设置集中的日志Writer
	//不管是我们的应用主动输出的日志，还是中间件，比如数据库等输出的日志，统一用Writer来管理，所以该Writer应该是一个公共的变量
	//创建文件，设置为writer
	//还要考虑因为模式的不同，日志的表现形式不同，在debug模式下，日志是直接输出到控制台，而release模式下，要记录日志
	setLoggerWriter()

	//初始化配置日志信息，例如格式等
	initLogger()

}

// 初始化日志信息
func initLogger() {
	//使用json模式记录,会返回一个log指针，需要保存起来，让别人调用
	logger = slog.New(slog.NewJSONHandler(logWriter, &slog.HandlerOptions{})) //初始化日志记录的格式

}

// logger
var logger *slog.Logger

func Logger() *slog.Logger {
	return logger
}

// 设置writer
func setLoggerWriter() {
	//逻辑，根据不同的mode，选择不同的writer
	switch gin.Mode() {
	case gin.ReleaseMode:
		//创建日志文件--以月来分割日志
		mouth := time.Now().Format("200601") //这是就是string类型

		//查找配置文件中的日志文件路径
		logfile := viper.GetString("app.log.path")
		logfile += fmt.Sprintf("/app-%s.log", mouth)

		//判断路径上的目录是否存在
		dir := filepath.Dir(logfile)
		//创建目录
		if err := os.MkdirAll(dir, 0666); err != nil {
			log.Println(err)
			return
		}

		if file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
			log.Println(err)
			return
		} else {
			logWriter = file
		}

	case gin.TestMode, gin.DebugMode:
		fallthrough
	default:
		logWriter = os.Stdout
	}
}

// 设置公共的Writer变量
var logWriter io.Writer //必须对该logWriter进行处理，不能被别人拿到了,小写的变量只能在本包中，utils包中可以被引用

func LogWriter() io.Writer {
	return logWriter
}
