package config

type Upload struct {
	SavePath  string `mapstructure:"save-path"`
	AccessUrl string `mapstructure:"access-url"`
}
