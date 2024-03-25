package repo

import (
	"vk_test/internal/database"
	"vk_test/internal/models"
)

type AnnouncementRepository interface {
	Store(models.Announcement) error
	Select() []models.Announcement
	SelectPage(offset, count int) []models.Announcement
	SelectById(id int) (models.Announcement, error)
	SelectWithFilter(filter database.SelectOptions) ([]models.Announcement, error)
}
