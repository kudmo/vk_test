package services

import (
	"net/http"
	"vk_test/internal/database"
	"vk_test/internal/database/interfaces"
	"vk_test/internal/models"
	"vk_test/internal/services/usecase/controller"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Interactor controller.UserInteractor
}

func NewUserController(sqlHandler interfaces.SqlHandler) *UserController {
	return &UserController{
		Interactor: controller.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) {
	u := models.User{}
	c.Bind(&u)
	controller.Interactor.Add(u)
	c.JSON(http.StatusCreated, u)
	return
}
