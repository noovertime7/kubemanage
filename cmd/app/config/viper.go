package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Binding(filePath string) error {
	v := viper.New()
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	// 把读取到的配置信息反序列化到 SysConfig 变量中
	if err := v.Unmarshal(&SysConfig); err != nil {
		return fmt.Errorf("config Unmarshal failed, err:%v\n", err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed,sys config reload")
		if err := viper.Unmarshal(&SysConfig); err != nil {
			fmt.Printf("config file changed,viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return nil
}
