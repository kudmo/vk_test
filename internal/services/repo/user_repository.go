package repo

import (
	"vk_test/internal/models"
)

type UserRepository interface {
	Store(models.User) error
	SelectById(id int) (models.User, error)
	SelectByLogin(login string) (models.User, error)
}
