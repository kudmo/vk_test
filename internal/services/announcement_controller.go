package services

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"vk_test/internal/auth"
	"vk_test/internal/database"
	"vk_test/internal/database/interfaces"
	"vk_test/internal/models"
	"vk_test/internal/services/repo"

	"github.com/labstack/echo/v4"
)

type AnnouncementController struct {
	AnnouncementRepository repo.AnnouncementRepository
}

func NewAnnouncementController(sqlHandler interfaces.SqlHandler) *AnnouncementController {
	return &AnnouncementController{
		AnnouncementRepository: &database.AnnouncementRepository{
			SqlHandler: sqlHandler,
		},
	}
}

func (controller *AnnouncementController) Get(c echo.Context) error {
	type Info struct {
		Title        string `json:"title"`
		Text         string `json:"text"`
		Price        int    `json:"price"`
		Image        string `json:"image_ref"`
		Author       string `json:"author"`
		Me_is_author bool   `json:"is_author"`
		Date         string `json:"date_creation"`
	}

	filter := database.SelectOptions{}
	echo.QueryParamsBinder(c)
	if err := c.Bind(&filter); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if filter.PageNumber < 0 {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if filter.OrderBy != "" && filter.OrderBy != "created_at" && filter.OrderBy != "price" {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if filter.PageNumber == 0 {
		filter.PageNumber = 1
	}

	userId, valid := auth.TokenGetUserId(c)
	if !valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid or expired jwt"})
	}

	ads, err := controller.AnnouncementRepository.SelectWithFilter(filter)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	var res []Info

	for i, _ := range ads {
		res = append(res, Info{})
		res[i].Title = ads[i].Title
		res[i].Text = ads[i].Text
		res[i].Price = ads[i].Price
		res[i].Author = ads[i].Client.Login
		res[i].Date = ads[i].CreatedAt.Format(time.DateOnly)
		res[i].Image = ads[i].ImageRef

		res[i].Me_is_author = ads[i].ClientId == userId
	}
	return c.JSON(http.StatusOK, res)
}

func (controller *AnnouncementController) Create(c echo.Context) error {
	data := models.Announcement{}

	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad data")
	}
	userId, _ := auth.TokenGetUserId(c)
	fmt.Println(userId)
	data.ClientId = userId
	err := controller.AnnouncementRepository.Store(data)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}

	return c.JSON(http.StatusCreated, data)
}
