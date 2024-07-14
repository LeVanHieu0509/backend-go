package controller

import (
	"fmt"
	"net/http"

	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
	"github.com/LeVanHieu0509/backend-go/pkg/ultis"
	ty "github.com/LeVanHieu0509/backend-go/types_custom"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: service.NewAuthService(),
	}
}

func (auth *AuthController) Login(ctx *gin.Context) {
	var requestBody ty.LoginReq

	err := ctx.ShouldBindJSON(&requestBody)
	ultis.HandleShouldBindJSONErr(err, "Binding Failed login")

	// Print the parsed request body
	fmt.Println("Request Body:", requestBody)
	response.SuccessResponse(ctx, 20001, auth.authService.Login(requestBody))
	ctx.JSON(http.StatusOK, gin.H{"user": requestBody})

}
func (auth *AuthController) SignUp() {

}

func (auth *AuthController) Logout() {

}

func (auth *AuthController) RefreshToken() {

}

func (auth *AuthController) ChangePass() {

}

func (auth *AuthController) DeleteUser() {

}

func (auth *AuthController) ResetPass() {

}

func (auth *AuthController) GetUsers() {

}
