package main

import (
	"chat/db"
	"chat/handlers"
	"chat/repository"
	"chat/router"
	"chat/service"
	"chat/settings"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
)

func main() {

	ctx := context.Background()

	if err := settings.LoadConfig(ctx); err != nil {
		log.Fatalf("Error loading settings: %v", err)
	}

	set := settings.GetSettings()
	db := db.ConnectDb()
	userRepo := repository.NewSqlUserRepository(db.GetDb())

	userService := service.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())

	e.GET("/app-health/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"health": "Ok",
		})
	})

	router.InitRouter(e, userHandler)

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	// Server
	if err := e.StartH2CServer(fmt.Sprintf(":%d", set.Port), s); err != nil {
		log.Fatal(err)
	}

}
