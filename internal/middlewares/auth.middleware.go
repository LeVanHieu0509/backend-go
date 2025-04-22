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
		// Trích xuất JWT token từ header của request
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

		// 1. context được sử dụng để gắn giá trị (ở đây là claims.Subject, chứa thông tin về UUID của người dùng)
		// 2. Điều này cho phép giá trị này được truyền qua các bước
		// xử lý tiếp theo trong pipeline của Gin mà không cần phải truyền trực tiếp qua các tham số của hàm.
		context := context.WithValue(ctx.Request.Context(), "subjectUUID", claims.Subject)

		// Lưu trữ giá trị claims.Subject vào context của request để có thể sử dụng trong các bước xử lý tiếp theo.
		// Sau khi gắn giá trị vào context, bạn phải cập nhật lại ctx.Request để chứa context mới:
		ctx.Request = ctx.Request.WithContext(context)
		ctx.Next()
	}
}

/*
	1. context cho phép bạn truyền các giá trị như claims.Subject (UUID của người dùng)
	qua nhiều hàm mà không cần phải thêm các tham số vào hàm.
	2. Khi subjectUUID đã được lưu vào context, bạn có thể dễ dàng truy xuất nó
	ở các bước tiếp theo mà không lo lắng về việc mất mát dữ liệu trong quá trình xử lý request.
*/
