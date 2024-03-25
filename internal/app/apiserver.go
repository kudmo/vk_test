package app

import (
	"fmt"
	"vk_test/internal/config"
	"vk_test/internal/models"
	"vk_test/internal/transport"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	state bool
}

func (server *Server) Start() {
	dbinit()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	transport.Init(e)

	e.Logger.Fatal(e.Start(config.ServerPort))
}

func dbinit() {

	db, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{})

	err = db.Migrator().CreateTable(models.User{})
	if err != nil {
		fmt.Print("User already exists")
	}

	err = db.Migrator().CreateTable(models.Announcement{})
	if err != nil {
		fmt.Print("Announcement already exists")
	}
}
