package usecase

import (
	"gorm-v2/repository"
	"gorm-v2/usecase/user"
)

type UseCase struct {
	User user.IUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User: user.New(repo),
	}
}
