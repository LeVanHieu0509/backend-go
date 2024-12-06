package user

import (
	"github.com/LeVanHieu0509/backend-go/internal/controller/account"
	"github.com/LeVanHieu0509/backend-go/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router

	// this is non dependency

	// ur := repo.NewUserRepo()
	// us := service.NewUserService(ur)
	// userHandleNonDependency := controller.NewUserController(us)

	// this is dependency => use pattern dependency injection => video 18  Wire Dependency Injection (Kiểm tra or nâng cấp)
	// Cho phép các module cấp cao tách biệt
	// userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)
		userRouterPublic.POST("/login", account.Login.Login)
		userRouterPublic.POST("/update_pass_register", account.Login.UpdatePasswordRegister)

		userRouterPublic.GET("/otp")

	}

	// Private router
	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthMiddleware()) //Limiter
	userRouterPrivate.Use()                             //Authentication
	userRouterPrivate.Use()                             //Permission

	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)
		userRouterPrivate.POST("/two-factor/verify", account.TwoFA.VerifyTwoFactorAuth)

	}
}
