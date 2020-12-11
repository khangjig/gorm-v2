package user

import (
	"gorm-v2/repository"
	"gorm-v2/repository/user"
)

type UseCase struct {
	UserRepo user.Repository
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		UserRepo: repo.User,
	}
}
