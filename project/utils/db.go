package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 初始化数据库连接 InitDB()
func InitDB() {
	//构建logger
	var logLevel gormLogger.LogLevel

	//根据应用的mode来控制 日志级别

	switch gin.Mode() {
	case gin.ReleaseMode:
		logLevel = gormLogger.Warn
	case gin.TestMode, gin.DebugMode:
		fallthrough
	default:
		//最高级别的日志输出（最详细）Info级别 最高级
		logLevel = gormLogger.Info
	}
	dbLogger := gormLogger.New(log.New(LogWriter(), "\n", log.LstdFlags), gormLogger.Config{
		SlowThreshold:             time.Second, //超过1s的 认为是慢查询
		Colorful:                  false,
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      false,
		LogLevel:                  logLevel,
	})

	//1、配置,是针对 gorm.Config{} 对象的处理
	conf := &gorm.Config{
		SkipDefaultTransaction:                   false,    //关闭默认事务
		DisableForeignKeyConstraintWhenMigrating: true,     //数据表迁移的时候禁用外键约束
		Logger:                                   dbLogger, //这里需要先构建
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, //数据表命名的策略
	}

	//2、创建db对象
	dsn := viper.GetString("db.dsn")
	fmt.Println(dsn)
	if dbNew, err := gorm.Open(mysql.Open(dsn), conf); err != nil {
		log.Fatalln(err)

	} else {
		db = dbNew
	}

}

// 建立全局的db对象，且封装之后，外部只能调用对象，不能修改对象，修改只能在本文件内进行
var db *gorm.DB

func DB() *gorm.DB {
	return db
}
