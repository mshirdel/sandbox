package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

var _runCmd = &cobra.Command{
	Use:   "run",
	Short: "run some code",
	RunE: func(cmd *cobra.Command, args []string) error {
		run()
		return nil
	},
}

func run() {
	fmt.Println(os.Hostname())
	e := echo.New()
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "test is ok")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
