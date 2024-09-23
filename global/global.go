package global

import (
	"database/sql"

	"github.com/LeVanHieu0509/backend-go/pkg/logger"
	"github.com/LeVanHieu0509/backend-go/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *gorm.DB
	Mdbc          *sql.DB
	Rdb           *redis.Client
	KafkaProducer *kafka.Writer
)

/*
	Config
	Mysql
	redis
*/
