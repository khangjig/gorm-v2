package user

import (
	"context"
	"fmt"
	"gorm-v2/model"
	"strings"
)

func (u *UseCase) GetAll(ctx context.Context, req *model.GetAllRequest) (*model.UsersResponse, error) {
	order := make([]string, 0)

	if req.Paginator.Limit < 0 {
		req.Paginator.Limit = -1
	}

	if req.Paginator.Page <= 0 {
		req.Paginator.Page = 1
	}

	if req.SortBy != "" {
		req.SortBy = model.SortType(strings.ToUpper(string(req.SortBy)))

		if req.SortBy != model.SortTypeASC && req.SortBy != model.SortTypeDESC {
			req.SortBy = model.SortTypeASC
		}
	}

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s %s", req.OrderBy, req.SortBy))
	}

	data, total := u.UserRepo.GetAll(ctx,
		[]model.Condition{},
		model.Paginator{
			Page:  req.Paginator.Page,
			Limit: req.Paginator.Limit,
		},
		order,
	)

	return &model.UsersResponse{
		Users: data,
		Page:  req.Paginator.Page,
		Limit: req.Paginator.Limit,
		Total: total,
	}, nil
}
