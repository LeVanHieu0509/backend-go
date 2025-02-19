package initialize

import (
	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/middlewares"
	"github.com/LeVanHieu0509/backend-go/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Nếu trong môi trường dev thì cần phải ghi log lại
	// r := gin.Default()
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// Khai báo theo sơ đồ Big team 1
	// r.Use() //Logger
	// r.Use() //Cross
	r.Use(middlewares.NewRateLimiter().GlobalRateLimiter()) //Limiter global
	r.GET("/ping/100", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong 100",
		})
	})

	r.Use(middlewares.NewRateLimiter().PublicAPIRateLimiter()) //Limiter global
	r.GET("/ping/80", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong 80",
		})
	})

	r.Use(middlewares.NewRateLimiter().UserAndPrivateAPIRateLimiter()) //Limiter global
	r.GET("/ping/50", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong 50",
		})
	})

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/check-status") //tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
		userRouter.InitTicketRouter(MainGroup)
	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)
	}

	return r
}
