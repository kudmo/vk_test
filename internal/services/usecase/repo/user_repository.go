package repo

import (
	"vk_test/internal/models"
)

type UserRepository interface {
	Store(models.User)
}
