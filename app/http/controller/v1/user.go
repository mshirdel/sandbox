package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}
