package database

import (
	"vk_test/internal/config"
	"vk_test/internal/database/interfaces"
	"vk_test/internal/models"

	"gorm.io/gorm"
)

type AnnouncementRepository struct {
	interfaces.SqlHandler
}

type SelectOptions struct {
	OrderBy    string `query:"order_by"`
	PageNumber int    `query:"page"`
}

func (db *AnnouncementRepository) Store(ad models.Announcement) error {
	return db.Create(&ad)
}

func (db *AnnouncementRepository) SelectById(id int) (models.Announcement, error) {
	var ad models.Announcement
	res := db.Where("id = ?", id).Find(&ad)
	if res.RowsAffected == 0 {
		return ad, gorm.ErrRecordNotFound
	}
	return ad, res.Error
}

func (db *AnnouncementRepository) Select() []models.Announcement {
	var ads []models.Announcement
	return ads
}

func (db *AnnouncementRepository) SelectPage(offset, count int) []models.Announcement {
	var ads []models.Announcement
	db.Offset(offset).Limit(count).Find(&ads)
	return ads
}
func (db *AnnouncementRepository) SelectWithFilter(filter SelectOptions) ([]models.Announcement, error) {
	var ads []models.Announcement
	var res *gorm.DB
	res = nil

	if filter.OrderBy == "" {
		filter.OrderBy = "created_at"
	}

	res = db.Preload("Client").Order(filter.OrderBy).Offset(config.PageSize * (filter.PageNumber - 1)).Limit(config.PageSize).Find(&ads)
	if res.Error != nil {
		return nil, res.Error
	}
	return ads, nil
}
