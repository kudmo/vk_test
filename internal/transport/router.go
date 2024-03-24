package transport

import (
	"vk_test/internal/database"
	controllers "vk_test/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
	e.Use(middleware.Recover())

	userController := controllers.NewUserController(database.NewSqlHandler())

	e.POST("/users", userController.Create)

	e.Logger.Fatal(e.Start(":1323"))
}
