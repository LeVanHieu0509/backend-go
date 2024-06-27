package middlewares

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		fmt.Printf("Before --> AuthMiddleware\n")
		token := ctx.GetHeader("Authorization")

		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrInvalidToken, "")
			ctx.Abort()
			return
		}
		ctx.Next()
		fmt.Println("Alter --> AuthMiddleware")
	}
}
