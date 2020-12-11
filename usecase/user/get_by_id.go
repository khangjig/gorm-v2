package user

import (
	"context"
	"errors"
	"gorm-v2/model"
	"gorm-v2/util/myerror"

	"gorm.io/gorm"
)

func (u *UseCase) GetByID(ctx context.Context, id int64) (*model.UserResponse, error) {
	resp := &model.UserResponse{}

	data, err := u.UserRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}

		return nil, myerror.ErrGetByID(err)
	}

	resp.User = data

	return resp, nil
}
