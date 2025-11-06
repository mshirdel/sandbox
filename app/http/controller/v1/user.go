package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetUserRequest struct {
	ID   uint   `param:"id" requied:"true"`
	Sort string `query:"sort" requied:"true"`
}

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetUser(ctx echo.Context) error {
	req := new(GetUserRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, req)
}
