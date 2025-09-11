package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (h *HomeController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, StringResponse{Status: 200, Data: "OK"})
}

type StringResponse struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}
