package controller

import (
	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
	"github.com/gin-gonic/gin"
)

// 1. sử dụng struct
// type UserController struct {
// 	userService *service.UserService
// }

// // 2. sử dụng con trỏ
// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// // controller --> service --> repo --> models --> dbs.
// func (uc *UserController) GetUserById(ctx *gin.Context) {
// 	a := uc.userService.GetInfoUser()
// 	fmt.Println(a)

// 	response.SuccessResponse(ctx, 20001, []string{"ok"})
// }

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")

	response.SuccessResponse(c, result, nil)
}
