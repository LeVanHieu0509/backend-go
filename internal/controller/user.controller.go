package controller

import (
	"fmt"
	"net/http"

	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/gin-gonic/gin"
)

// 1. sử dụng struct
type UserController struct {
	userService *service.UserService
}

// 2. sử dụng con trỏ
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller --> service --> repo --> models --> dbs.
func (uc *UserController) GetUserById(ctx *gin.Context) {
	a := uc.userService.GetInfoUser()
	fmt.Println(a)
	ctx.JSON(http.StatusOK, gin.H{
		"ok": 1,
	})
}
