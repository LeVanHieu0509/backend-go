// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/LeVanHieu0509/backend-go/internal/controller"
	"github.com/LeVanHieu0509/backend-go/internal/repo"
	"github.com/LeVanHieu0509/backend-go/internal/service"
)

// Injectors from user.wire.go:

func InitUserRouterHandler() (*controller.UserController, error) {
	iUserRepository := repo.NewUserRepository()
	iAuthRepository := repo.NewAuthRepository()
	iUserService := service.NewUserService(iUserRepository, iAuthRepository)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
