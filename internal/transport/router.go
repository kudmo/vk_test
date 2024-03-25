package transport

import (
	"vk_test/internal/auth"
	"vk_test/internal/config"
	"vk_test/internal/database"
	controllers "vk_test/internal/services"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
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
	announcementController := controllers.NewAnnouncementController(database.NewSqlHandler())

	userGroup := e.Group("/users")
	userGroup.POST("/registrate", userController.Create)
	userGroup.POST("/login", userController.Login)

	announcementGroup := e.Group("/announcement")
	announcementGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/announcement/all"
		},
		SigningKey: []byte(config.SecretKeyJWT),
	}))

	announcementGroup.POST("/create", announcementController.Create)
	announcementGroup.GET("/all", announcementController.Get)

	e.Logger.Fatal(e.Start(":1323"))
}
