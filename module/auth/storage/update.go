package authstorage

import (
	"context"
	common "go-base/commons"
	authmodel "go-base/module/auth/model"
)

func (s *sqlStore) Update(ctx context.Context, id uint, data *authmodel.User) (error) {
	if err := s.db.Where("id=?", id).Updates(&data).Error; err != nil {
		return  common.ErrDB(err)
	}
	return nil
}