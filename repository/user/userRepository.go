package user

import (
	"context"

	"gorm-v2/model"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetAll(ctx context.Context,
		conditions []model.Condition,
		paginator model.Paginator,
		order []string,
	) ([]model.User, int64)
	Create(ctx context.Context, obj *model.User) error
	ExecuteSQL(ctx context.Context, query string) error
}
