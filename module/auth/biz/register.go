package authbiz

import (
	"context"
	"errors"
	authmodel "go-base/module/auth/model"
)
type Hasher interface{
	HashPassword(password string) (string, error)
	CheckPasswordAndHash(password, hash string) bool
}
type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*authmodel.User, error)
	CreateUser(ctx context.Context, data *authmodel.User) error
}

type registerBusiness struct{
	registerStorage RegisterStorage
	hasher Hasher
}
func NewRegisterBusiness(registerStorage RegisterStorage,hasher Hasher) *registerBusiness {
	return &registerBusiness{registerStorage: registerStorage,hasher: hasher}
}
func (business *registerBusiness) Register(ctx context.Context, data *authmodel.UserCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return errors.New("user has been existed")
	}
	entity:=authmodel.User{
		Email:data.Email,
		PhoneNo: data.PhoneNo,
	}
	hashPassword,err:=business.hasher.HashPassword(data.Password)
	entity.Password=hashPassword
	if err!=nil{
		return err
	}
	if err := business.registerStorage.CreateUser(ctx, &entity); err != nil {
		return err
	}
	return nil
}