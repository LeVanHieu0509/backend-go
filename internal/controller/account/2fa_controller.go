package account

import (
	"github.com/LeVanHieu0509/backend-go/internal/model"
	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/LeVanHieu0509/backend-go/internal/utils/context"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
	"github.com/gin-gonic/gin"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct{}

// User Setup Two Factor Authentication documentation
// @Summary      Setup Two Factor Authentication
// @Description  Setup Two Factor Authentication
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization token"
// @Param        payload body model.SetupTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two-factor/setup [post]
func (c *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid set up two factor auth parameter")

		return
	}

	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "UserId is not invalid")
		return
	}

	params.UserId = uint32(userId)

	codeResult, err := service.UserLogin().SetupTwoFactorAuth(ctx, &params)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeResult, nil)

	//2. get userId from uuid(token) when login success
}

// User Verify Two Factor Authentication documentation
// @Summary      Verify Two Factor Authentication
// @Description  Verify Two Factor Authentication
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization token"
// @Param        payload body model.VerifyTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two-factor/verify [post]
func (c *sUser2FA) VerifyTwoFactorAuth(ctx *gin.Context) {
	var params model.VerifyTwoFactorAuthInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid set up two factor auth parameter")

		return
	}

	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "UserId is not invalid")
		return
	}

	params.UserId = uint32(userId)
	codeResult, err := service.UserLogin().VerifyTwoFactorAuth(ctx, &params)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeResult, nil)

}
