package dao

import (
	"fmt"
	"github.com/noovertime7/kubemanage/config"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	isInit bool
	Gorm   *gorm.DB
	err    error
)

func InitDB() {
	fmt.Println("开始初始化DB")
	if isInit {
		return
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//
	////打印sql语句
	//连接池最大允许的空闲连接数
	sqlDB, err := Gorm.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	//设置最大连接数
	sqlDB.SetMaxIdleConns(config.MaxOpenConns)
	//设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(config.MaxLifetime * time.Second)
	isInit = true
}
