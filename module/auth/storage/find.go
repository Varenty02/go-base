package authstorage

import (
	"context"
	authmodel "go-base/module/auth/model"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*authmodel.User, error) {
	db := s.db.Table(authmodel.User{}.TableName())
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var user authmodel.User
	if err := db.Where(conditions).First(&user).Error; err != nil {
		return nil,err
	}
	return &user, nil
}