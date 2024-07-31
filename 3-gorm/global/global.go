package global

import (
	"go-learning/3-gorm/config"
	"gorm.io/gorm"
)

var AppConfig *config.AppConfig
var DB *gorm.DB
