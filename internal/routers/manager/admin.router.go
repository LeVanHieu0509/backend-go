package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// Public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")

	}
	// Private router
	adminRouterPrivate := Router.Group("/admin/user")
	adminRouterPrivate.Use() //Limiter
	adminRouterPrivate.Use() //Authentication
	adminRouterPrivate.Use() //Permission

	{
		adminRouterPrivate.POST("/active_user")

	}
}
