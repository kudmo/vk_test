package services

import (
	"log"
	"net/http"
	"regexp"
	"vk_test/internal/auth"
	"vk_test/internal/database"
	"vk_test/internal/database/interfaces"
	"vk_test/internal/models"
	"vk_test/internal/services/repo"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserRepository repo.UserRepository
}

func NewUserController(sqlHandler interfaces.SqlHandler) *UserController {
	return &UserController{
		UserRepository: &database.UserRepository{
			SqlHandler: sqlHandler,
		},
	}
}

func (interactor *UserController) checkAuth(login, password string) (int, error) {
	user, err := interactor.UserRepository.SelectByLogin(login)

	if err != nil || user.Login != login || user.Password != auth.HashPassword(password) {
		return 0, echo.ErrUnauthorized
	}

	return user.Id, nil
}

func (controller *UserController) Create(c echo.Context) error {
	type UserData struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	data := UserData{}
	password_regex, _ := regexp.Compile("^[a-zA-Z0-9_!@#$%^&*]{7,}$")
	login_regex, _ := regexp.Compile("^[a-zA-Z0-9_]{6,}$")
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad data")
	}

	login := data.Login
	if !login_regex.MatchString(login) {
		return c.String(http.StatusBadRequest, "Incorrect login")
	}
	password := data.Password
	if !password_regex.MatchString(password) {
		return c.String(http.StatusBadRequest, "Incorrect password")
	}
	hashedPassword := auth.HashPassword(password)

	u := models.User{Login: login, Password: hashedPassword}

	if controller.UserRepository.Store(u) != nil {
		return c.String(http.StatusBadRequest, "login already exist")
	}
	return c.JSON(http.StatusCreated, u)
}

func (controller *UserController) Login(c echo.Context) error {
	type UserData struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	data := UserData{}

	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad data")
	}

	userId, err := controller.checkAuth(data.Login, data.Password)
	if err != nil {
		return c.String(http.StatusUnauthorized, "wrong login or password")
	}

	atoken, err := auth.CalculateToken(userId)

	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": atoken,
	})
}
