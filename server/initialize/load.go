package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"mall.com/global"
)

// LoadConfig 加载配置文件
func LoadConfig() {
	viper.AddConfigPath("./server/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error resources file: %w \n", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct %w \n", err))
	}
}

