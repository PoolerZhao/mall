package config

type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Upload Upload `mapstructure:"upload"`
	Jwt    Jwt    `mapstructure:"jwt"`
}
