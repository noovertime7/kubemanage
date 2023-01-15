package config

var SysConfig *Config

type Config struct {
	Default DefaultOptions `mapstructure:"default"`
	Mysql   MysqlOptions   `mapstructure:"mysql"`
	CMDB    CMDBOptions    `mapstructure:"cmdb"`
	Log     LogConfig      `mapstructure:"log"`
}

type DefaultOptions struct {
	PodLogTailLine       string `mapstructure:"podLogTailLine"`
	ListenAddr           string `mapstructure:"listenAddr"`
	WebSocketListenAddr  string `mapstructure:"webSocketListenAddr"`
	JWTSecret            string `mapstructure:"JWTSecret"`
	ExpireTime           int64  `mapstructure:"expireTime"`
	KubernetesConfigFile string `mapstructure:"kubernetesConfigFile"`
}

type CMDBOptions struct {
	HostCheck HostCheck `mapstructure:"hostCheck"`
}

type HostCheck struct {
	HostCheckEnable   bool `mapstructure:"hostCheckEnable"`
	HostCheckDuration int  `mapstructure:"hostCheckDuration"`
	HostCheckTimeout  int  `mapstructure:"hostCheckTimeout"`
}

type MysqlOptions struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Port         string `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxLifetime  int    `mapstructure:"maxLifetime"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
