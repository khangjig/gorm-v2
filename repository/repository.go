package repository

import (
	"context"
	"gorm-v2/repository/user"

	"gorm.io/gorm"
)

type Repository struct {
	User user.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		User: user.NewPG(getClient),
	}
}
