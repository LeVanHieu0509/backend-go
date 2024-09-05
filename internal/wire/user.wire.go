//go:build-wireinject

package wire

import (
	"github.com/LeVanHieu0509/backend-go/internal/controller"
	"github.com/LeVanHieu0509/backend-go/internal/repo"
	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandlerInjection() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
