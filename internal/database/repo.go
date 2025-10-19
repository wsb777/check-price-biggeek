package database

import (
	"context"

	"gorm.io/gorm"
)

type Repo interface {
	CreateUser(ctx context.Context, chatId int64, userId int64, username string) error
	UpdateLinks()
	DeleteLinks()
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateUser(ctx context.Context, chatId int64, userId int64, username string) error {
	err := gorm.G[User](r.db).Create(ctx,
		&User{
			UserID:   userId,
			ChatID:   chatId,
			Username: username,
		})
	return err
}

func (r *repo) UpdateLinks() {

}

func (r *repo) DeleteLinks() {

}
