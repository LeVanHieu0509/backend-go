package middlewares

import (
	"context"
	"log"

	"github.com/LeVanHieu0509/backend-go/internal/utils/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// get the request url path
		uri := ctx.Request.URL.Path
		log.Println("URI Request: ", uri)

		//check headers authentication
		jwtToken, err := auth.ExtractBearerToken(ctx)
		log.Println("Token exact:: ", jwtToken, err)
		if !err {
			ctx.AbortWithStatusJSON(403, gin.H{"code": 40003, "err": "UnAuthorized", "description": ""})
			return
		}

		//validate jwt token by exact
		claims, valid := auth.VerifyTokenSubject(jwtToken)
		log.Println("claims, valid", claims, valid)
		if valid != nil {
			ctx.AbortWithStatusJSON(403, gin.H{"code": 40003, "err": "Invalid Token", "description": ""})
			return
		}

		log.Println("Claims::UUID::", claims.Subject) //
		context := context.WithValue(ctx.Request.Context(), "subjectUUID", claims.Subject)

		ctx.Request = ctx.Request.WithContext(context)
		ctx.Next()
	}
}
