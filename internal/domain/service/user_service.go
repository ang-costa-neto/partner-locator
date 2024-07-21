package service

import (
	"partner-locator/internal/domain/model"
	"partner-locator/internal/infrastructure/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo: repo}
}

func (s UserService) GetUsers() ([]model.User, error) {
	return s.repo.GetUsers()
}

func (s UserService) CreateUser(user model.User) error {
	return s.repo.CreateUser(user)
}

func (s UserService) UpdateUser(user model.User) error {
	return s.repo.UpdateUser(user)
}

func (s UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
