package user

import (
	"context"
	"gorm-v2/util/myerror"
)

func (u *UseCase) DeleteByID(ctx context.Context, id int64) error {
	err := u.UserRepo.Delete(ctx, id, false)
	if err != nil {
		return myerror.ErrDeleteByID(err)
	}

	return nil
}
