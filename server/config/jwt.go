package config

type Jwt struct {
	SigningKey string `mapstructure:"signing-key"`
}
