package account

import (
	"github.com/LeVanHieu0509/backend-go/internal/model"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
	"github.com/gin-gonic/gin"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct{}

func (c *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid set up two factor auth parameter")

		return
	}

	//2. get userId from uuid(token) when login success
}
