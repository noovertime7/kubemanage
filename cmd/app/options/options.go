package options

import (
	"fmt"
	"github.com/noovertime7/kubemanage/cmd/app/config"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

const (
	defaultConfigFile = "/etc/gopixiu/config.yaml"
)

// Options has all the params needed to run a pixiu
type Options struct {
	// The default values.
	ComponentConfig config.Config
	DB              *gorm.DB
	Factory         dao.ShareDaoFactory // 数据库接口
	// ConfigFile is the location of the pixiu server's configuration file.
	ConfigFile string
}

func NewOptions() (*Options, error) {
	return &Options{
		ConfigFile: defaultConfigFile,
	}, nil
}

func (o *Options) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.ConfigFile, "configfile", "", "The location of the gopixiu configuration file")
}

// Complete completes all the required options
func (o *Options) Complete() error {
	// 配置文件优先级: 默认配置，环境变量，命令行
	if len(o.ConfigFile) == 0 {
		// Try to read config file path from env.
		if cfgFile := os.Getenv("ConfigFile"); cfgFile != "" {
			o.ConfigFile = cfgFile
		} else {
			o.ConfigFile = defaultConfigFile
		}
	}

	c := config.New()
	c.SetConfigFile(o.ConfigFile)
	c.SetConfigType("yaml")
	if err := c.Binding(&o.ComponentConfig); err != nil {
		return err
	}

	// 注册依赖组件
	if err := o.register(); err != nil {
		return err
	}
	return nil
}

func (o *Options) register() error {
	if err := o.registerDatabase(); err != nil {
		return err
	}
	return nil
}

func (o *Options) registerDatabase() error {
	sqlConfig := o.ComponentConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConfig.User,
		sqlConfig.Password,
		sqlConfig.Host,
		sqlConfig.Port,
		sqlConfig.Name)

	var err error
	if o.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}
	// 设置数据库连接池
	sqlDB, err := o.DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(o.ComponentConfig.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(o.ComponentConfig.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(o.ComponentConfig.Mysql.MaxLifetime) * time.Second)
	o.Factory = dao.NewShareDaoFactory(o.DB)
	return nil
}
