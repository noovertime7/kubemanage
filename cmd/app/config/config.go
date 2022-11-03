package config

var SysConfig *Config

type Config struct {
	Default DefaultOptions `yaml:"default"`
	Mysql   MysqlOptions   `yaml:"mysql"`
}

type DefaultOptions struct {
	PodLogTailLine       string `yaml:"podLogTailLine"`
	ListenAddr           string `yaml:"listenAddr"`
	WebSocketListenAddr  string `yaml:"webSocketListenAddr"`
	JWTSecret            string `yaml:"JWTSecret"`
	ExpireTime           int64  `yaml:"expireTime"`
	KubernetesConfigFile string `yaml:"kubernetesConfigFile"`
}

type MysqlOptions struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         string `yaml:"port"`
	Name         string `yaml:"name"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxLifetime  int    `yaml:"maxLifetime"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
}
