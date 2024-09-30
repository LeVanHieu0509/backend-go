package initialize

import (
	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/database"
	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/LeVanHieu0509/backend-go/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)

	// User service Interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
}
