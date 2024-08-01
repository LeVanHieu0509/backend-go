package user

import (
	"github.com/LeVanHieu0509/backend-go/internal/controller"
	"github.com/LeVanHieu0509/backend-go/internal/repo"
	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router

	// this is non dependency
	ur := repo.NewUserRepo()
	us := service.NewUserService(ur)
	userHandleNonDependency := controller.NewUserController(us)

	// this is dependency => use pattern dependency injection => video 18

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userHandleNonDependency.Register)
		userRouterPublic.GET("/otp")

	}

	// Private router
	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use() //Limiter
	userRouterPrivate.Use() //Authentication
	userRouterPrivate.Use() //Permission

	{
		userRouterPrivate.GET("/get_info")

	}
}
