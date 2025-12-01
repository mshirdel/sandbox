package v1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type GetUserRequest struct {
	ID   uint   `param:"id" requied:"true"`
	Sort string `query:"sort" requied:"true"`
}

type SaveMessageRequest struct {
	Title   string `json:"title" requied:"true"`
	Message string `json:"message" requied:"true"`
}

type SaveMessageResponse struct {
	ID uint `json:"id"`
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

func (u *UserController) SaveMessage(ctx echo.Context) error {
	req := new(SaveMessageRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	logrus.Infof("title: %s \t msg: %s \n", req.Title, req.Message)

	return ctx.JSON(http.StatusOK, SaveMessageResponse{ID: 1010})
}

func Add(n int) error {
	if n == 10 {
		return fmt.Errorf("number [%d] is ten", n)
	}

	return nil
}
