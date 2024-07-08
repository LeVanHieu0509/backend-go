package global

import (
	"github.com/LeVanHieu0509/backend-go/pkg/logger"
	"github.com/LeVanHieu0509/backend-go/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

/*
	Config
	Mysql
	redis
*/
