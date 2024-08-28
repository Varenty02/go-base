package authstorage

import (
	"context"
	authmodel "go-base/module/auth/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *authmodel.User) error {
	//mở 1 transactiona
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}