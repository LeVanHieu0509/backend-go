package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> AA")
		ctx.Next()
		fmt.Println("Alter --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> BB")
		ctx.Next()
		fmt.Println("Alter --> BB")
	}
}

func CC(ctx *gin.Context) {
	fmt.Println("Before --> CC")
	ctx.Next()
	fmt.Println("Alter --> CC")
}

// func InitRouterSmall() *gin.Engine {
// 	r := gin.Default() //func để tạo instance mặc định
// 	// r.Use(middlewares.AuthMiddleware(), AA(), BB(), CC)

// 	v1 := r.Group("v1/2024")
// 	{
// 		v1.GET("/ping", Pong)
// 		v1.GET("/user/1", c.NewUserController().GetUserById)
// 		v1.PATCH("/ping", Pong)
// 		v1.DELETE("/ping", Pong)
// 		v1.HEAD("/ping", Pong)
// 		v1.OPTIONS("/ping", Pong)
// 	}

// 	v2 := r.Group("v2/2024")
// 	{
// 		v2.GET("/ping", Pong)
// 		v2.PUT("/ping", Pong)
// 		v2.PATCH("/ping", Pong)
// 		v2.DELETE("/ping", Pong)
// 		v2.HEAD("/ping", Pong)
// 		v2.OPTIONS("/ping", Pong)
// 	}

// 	v3 := r.Group("v3/application")
// 	{
// 		v3.POST("/login", c.NewAuthController().Login)

// 	}

// 	return r
// }

// func Pong(ctx *gin.Context) { //ctx: xử lý request và response
// 	//Json trả về client format JSON
// 	// gin.H: map string {key: value}

// 	name := ctx.Param("name")
// 	age := ctx.DefaultQuery("age", "hieu") // If age not has, set age -> hieu
// 	uid := ctx.Query("uid")                //default
// 	fmt.Printf("My HANDLE\n")
// 	// Sử dụng Logger
// 	global.Logger.Info("Application started, trace-id")

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "pong" + name,
// 		"uid":     uid,
// 		"age":     age,
// 		"users":   []string{"cr07", "m10"},
// 	})
// }
