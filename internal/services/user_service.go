package services

import (
	"github.com/wsb777/check-price-biggeek/internal/database"
)

type UserService interface {
	RegisterUser(string) error
}

type userService struct {
	repo *database.Repo
}

func NewUserService(repo *database.Repo) UserService {
	return &userService{
		repo: repo,
	}
}

func (s userService) RegisterUser(string) error {

	return nil
}
