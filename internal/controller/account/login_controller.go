package account

import (
	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/model"
	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// manager controller login user
var Login = new(cUserLogin)

type cUserLogin struct {
}

// User Login documentation
// @Summary      User Login
// @Description  get user login
// @Tags         User Login
// @Accept       json
// @Produce      json
// @Param        payload body model.LoginInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/login [post]
func (c *cUserLogin) Login(ctx *gin.Context) {
	// User login
	var params model.LoginInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeRs, dataRs, err := service.UserLogin().Login(ctx, &params)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}

	response.SuccessResponse(ctx, codeRs, dataRs)
}

// User Registration documentation
// @Summary      Show an account
// @Description  get user register
// @Tags         User Registration
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/register [post]
func (c *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeStatus, err := service.UserLogin().Register(ctx, &params)

	if err != nil {
		global.Logger.Error("Error registration user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}

// User VerifyOTP documentation
// @Summary      Show an account
// @Description  user VerifyOTP
// @Tags         User VerifyOTP
// @Accept       json
// @Produce      json
// @Param        payload body model.VerifyInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/verify_account [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())

		return
	}
	result, err := service.UserLogin().VerifyOTP(ctx, &params)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidOtp, err.Error())
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}

// User UpdatePasswordRegister documentation
// @Summary      Show an account
// @Description  user UpdatePasswordRegister
// @Tags         User UpdatePasswordRegister
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateUserPasswordInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/update_pass_register [post]
func (c *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params model.UpdateUserPasswordInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())

		return
	}
	result, err := service.UserLogin().UpdatePassword(ctx, params.UserToken, params.UserPassword)

	if err != nil {
		response.ErrorResponse(ctx, result, err.Error())
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
