package services

import (
	"crypto/sha256"
	"encoding/hex"
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

func (controller *UserController) Create(c echo.Context) error {
	type UserData struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	data := UserData{}

	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad data")
	}

	login := data.Login
	password := data.Password
	sum := sha256.Sum256([]byte(password))
	hashedPassword := hex.EncodeToString(sum[:])

	u := models.User{Login: login, Password: hashedPassword}

	if controller.Interactor.Add(u) != nil {
		return c.String(http.StatusBadRequest, "login already exist")
	}
	return c.JSON(http.StatusCreated, u)
}
