package user

import (
	"context"
	"gorm-v2/model"
)

type IUseCase interface {
	GetByID(ctx context.Context, id int64) (*model.UserResponse, error)
	GetAll(ctx context.Context, req *model.GetAllRequest) (*model.UsersResponse, error)
	DeleteByID(ctx context.Context, id int64) error
}
