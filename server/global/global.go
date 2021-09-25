package global

import (
	"gorm.io/gorm"
	"mall.com/config"
)

var (
	Config config.Config
	Db    *gorm.DB
)

