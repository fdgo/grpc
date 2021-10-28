package main

type JwtConfig struct {
	Key string `mapstructure:"key" yaml:"key"`
}

type SmsConfig struct {
	Key     string `mapstructure:"key" yaml:"key"`
	Secretc string `mapstructure:"secretc" yaml:"secretc"`
	Expire  string `mapstructure:"expire" yaml:"expire"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port int    `mapstructure:"port" yaml:"port"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port int    `mapstructure:"port" yaml:"port"`
}

type ServerConfig struct {
	JwtConfig  JwtConfig    `mapstructure:"jwt" yaml:"jwt"`
	SmsInfo    SmsConfig    `mapstructure:"sms" yaml:"sms"`
	RedisInfo  RedisConfig  `mapstructure:"redis" yaml:"redis"`
	ConsulInfo ConsulConfig `mapstructure:"consul" yaml:"consul"`
}

//**************************************************************************
type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
