package router

import (
	"chat/handlers"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, userHandler handlers.UserHandler) {

	e.POST("/users/", userHandler.CreateUser)

}
