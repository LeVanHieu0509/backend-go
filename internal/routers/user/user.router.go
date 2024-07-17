package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
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
