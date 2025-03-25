package initialize

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitPrometheus()
	InitLogger()
	InitMySql()
	InitMySqlC()
	// InitRedisSentinel()
	InitRedis()
	InitKafka()
	InitServiceInterface()

	r := InitRouter()

	return r
}
