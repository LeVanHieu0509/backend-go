package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsRouter() *gin.Engine {
	r := gin.Default() //func để tạo instance mặc định

	v1 := r.Group("v1/2024")
	{
		v1.GET("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
	}

	v2 := r.Group("v2/2024")
	{
		v2.GET("/ping", Pong)
		v2.PUT("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.HEAD("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
	}

	return r
}

func Pong(ctx *gin.Context) { //ctx: xử lý request và response
	//Json trả về client format JSON
	// gin.H: map string {key: value}

	name := ctx.Param("name")
	age := ctx.DefaultQuery("age", "hieu") // If age not has, set age -> hieu
	uid := ctx.Query("uid")                //default

	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
		"age":     age,
		"users":   []string{"cr07", "m10"},
	})
}