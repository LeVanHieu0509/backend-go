package initialize

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/global"
)

func Run() {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)

	InitLogger()
	InitMySql()
	InitMySqlC()
	InitServiceInterface()
	InitRedis()
	InitKafka()
	r := InitRouter()
	r.Run(":8001")
}
