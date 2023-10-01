package handlers

import (
	"chat/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return UserHandler{service: s}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	u := new(service.CreateUserReq)
	err := (&echo.DefaultBinder{}).BindBody(c, u)
	if err != nil {
		return echo.NewHTTPError(echo.ErrUnprocessableEntity.Code, fmt.Sprintf("Error processing body:%v", err))
	}

	res, err := h.service.CreateUser(c.Request().Context(), *u)

	switch err {
	case service.ErrorUserAlreadyCreated:
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	case nil:
		return c.JSON(http.StatusOK, res)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

}
