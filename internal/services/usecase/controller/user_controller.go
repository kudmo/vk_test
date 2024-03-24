package controller

import (
	"vk_test/internal/models"
	"vk_test/internal/services/usecase/repo"
)

type UserInteractor struct {
	UserRepository repo.UserRepository
}

func (interactor *UserInteractor) Add(u models.User) error {
	return interactor.UserRepository.Store(u)
}
