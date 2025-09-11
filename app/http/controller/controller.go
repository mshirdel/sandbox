package controller

import (
	"fmt"
	"math"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mshirdel/sandbox/app"
)

type Controller struct {
	app            *app.Application
	homeController *HomeController
}

func NewController(app *app.Application) *Controller {
	return &Controller{
		app:            app,
		homeController: NewHomeController(),
	}
}

func (c *Controller) Routes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.GET("/", c.homeController.Index)

	printRoutes(e.Routes())

	return e
}

func printRoutes(routes []*echo.Route) {
	var maxMethodLength, maxPathLength float64
	for _, r := range routes {
		maxMethodLength = math.Max(maxMethodLength, float64(len(r.Method)))
		maxPathLength = math.Max(maxPathLength, float64(len(r.Path)))
	}

	fmt.Printf("\nRegistered http routes:\n")

	for _, r := range routes {
		// do not print middlewares
		if strings.HasPrefix(r.Name, "github.com/labstack/echo") {
			continue
		}

		fmt.Printf("%-*v %-*v --> %v\n", int(maxMethodLength), r.Method, int(maxPathLength), r.Path, r.Name)
	}
}
