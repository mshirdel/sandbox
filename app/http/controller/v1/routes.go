package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/mshirdel/sandbox/app"
)

type Router struct {
	app            *app.Application
	userController *UserController
}

func New(a *app.Application) *Router {
	return &Router{
		app:            a,
		userController: NewUserController(),
	}
}

func (r *Router) Routes(g *echo.Group) {
	user := g.Group("/users")
	{
		user.GET("", r.userController.GetUser)
	}
}
