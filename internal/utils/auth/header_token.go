package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractBearerToken(c *gin.Context) (string, bool) {
	// Lấy giá trị của header Authorization từ yêu cầu HTTP.
	authHeader := c.GetHeader("Authorization")

	if strings.HasPrefix(authHeader, "Bearer") {
		// Loại bỏ chuỗi "Bearer" khỏi đầu header.
		jwtToken := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		return jwtToken, true
	}
	return "", false
}
