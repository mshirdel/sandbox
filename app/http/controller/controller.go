package controller

import (
	"fmt"
	"math"
	"strings"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mshirdel/sandbox/app"
	v1 "github.com/mshirdel/sandbox/app/http/controller/v1"
)

type Controller struct {
	app *app.Application
	v1  *v1.Router
}

func NewController(app *app.Application) *Controller {
	return &Controller{
		app: app,
		v1:  v1.New(app),
	}
}

func (c *Controller) Routes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(echoprometheus.NewMiddleware("sandbox-app"))
	trace := jaegertracing.New(e, nil)
	defer trace.Close()

	e.GET("/metrics", echoprometheus.NewHandler())

	c.v1.Routes(e.Group("/v1"))
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
