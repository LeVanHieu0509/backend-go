package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router
	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.POST("/login")

	}
	// Private router
	userRouterPrivate := Router.Group("/admin/user")
	userRouterPrivate.Use() //Limiter
	userRouterPrivate.Use() //Authentication
	userRouterPrivate.Use() //Permission

	{
		userRouterPrivate.POST("/admin/active_user")

	}
}
