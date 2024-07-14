package global

import (
	"github.com/LeVanHieu0509/backend-go/pkg/logger"
	"github.com/LeVanHieu0509/backend-go/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)

/*
	Config
	Mysql
	redis
*/
