package database

import (
	"vk_test/internal/database/interfaces"
	"vk_test/internal/models"
)

type UserRepository struct {
	interfaces.SqlHandler
}

func (db *UserRepository) Store(u models.User) error {
	return db.Create(&u)
}
