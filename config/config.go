package config

import (
	"github.com/spf13/viper"
	"github.com/wonderivan/logger"
	_ "gopkg.in/yaml.v2"
	"os"
	"time"
)

var (
	ListenAddr          = SystemConf.ListenAddr
	PodLogTailLine      = SystemConf.PodLogTailLine
	DBUser              = SystemConf.Database.User
	DBPassword          = SystemConf.Database.Password
	DBHost              = SystemConf.Database.Host
	DBPort              = SystemConf.Database.Port
	DBName              = SystemConf.Database.DBName
	MaxOpenConns        = SystemConf.Database.MaxOpenConns
	MaxLifetime         = SystemConf.MaxLifetime * time.Second
	JWTSecret           = SystemConf.JWTSecret
	ExpireTime          = 1
	KubeConfigFile      = "C:\\Users\\18495\\.kube\\config"
	WebSocketListenAddr = ":9091"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		logger.Error(err)
		return
	}
	config := viper.New()
	config.SetConfigType("yaml")           //设置文件的类型
	config.AddConfigPath(path + "/config") //设置读取的文件路径
	config.SetConfigName("config")         //设置读取的文件名
	if err := config.ReadInConfig(); err != nil {
		logger.Fatal("读取配置文件失败", err)
		return
	}
	if err = config.Unmarshal(SystemConf); err != nil {
		logger.Fatal("配置文件解析失败", err)
		return
	}
}

var SystemConf = new(System)

type Database struct {
	Host         string `yaml:"host"`
	Password     string `yaml:"password"`
	User         string `yaml:"user"`
	Port         string `yaml:"port"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxLifetime  int    `yaml:"max_lifetime"`
	DBName       string `yaml:"db_name"`
}

type KubeManage struct {
	Database            Database `yaml:"Database"`
	PodLogTailLine      int      `yaml:"pod_log_tail_line"`
	ListenAddr          string   `yaml:"listen_addr"`
	WebSocketListenAddr string   `yaml:"web_socket_listen_addr"`
	KubeConfigFile      string   `yaml:"kube_config_file"`
	JWTSecret           string   `yaml:"jwt_secret"`
	ExpireTime          int      `yaml:"expire_time"`
}

type System struct {
	KubeManage
}
