package routers

import (
	"github.com/LeVanHieu0509/backend-go/internal/routers/manager"
	"github.com/LeVanHieu0509/backend-go/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
