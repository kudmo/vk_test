package database

import (
	"vk_test/internal/database/interfaces"
	"vk_test/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	interfaces.SqlHandler
}

func (db *UserRepository) Store(u models.User) error {
	return db.Create(&u)
}

func (db *UserRepository) SelectByLogin(login string) (models.User, error) {
	var user models.User
	res := db.Where("login = ?", login).Find(&user)
	if res.RowsAffected == 0 {
		return user, gorm.ErrRecordNotFound
	}
	return user, res.Error
}

func (db *UserRepository) SelectById(id int) (models.User, error) {
	var user models.User
	res := db.Where("id = ?", id).Find(&user)
	if res.RowsAffected == 0 {
		return user, gorm.ErrRecordNotFound
	}
	return user, res.Error
}
