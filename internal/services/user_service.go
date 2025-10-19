package services

import (
	"context"
	"log"

	"github.com/wsb777/check-price-biggeek/internal/database"
)

type UserService interface {
	RegisterUser(ctx context.Context, chatId int64, userID int64, username string) error
}

type userService struct {
	repo database.Repo
}

func NewUserService(repo database.Repo) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) RegisterUser(ctx context.Context, chatId int64, userID int64, username string) error {
	err := s.repo.CreateUser(ctx, chatId, userID, username)
	if err != nil {
		log.Fatalf("[ERROR] RegisterUser: %s", err)
		return err
	}
	return nil
}
