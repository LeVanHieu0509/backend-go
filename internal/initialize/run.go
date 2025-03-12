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

	InitLogger()
	InitMySql()
	InitMySqlC()
	InitServiceInterface()
	// InitRedis()
	InitRedisSentinel()
	// InitKafka()
	r := InitRouter()

	return r
}
