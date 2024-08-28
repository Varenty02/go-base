package authbiz

import (
	"context"
	common "go-base/commons"
	"go-base/component/jwt"
	authmodel "go-base/module/auth/model"
	"log"
	"time"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*authmodel.User, error)
	Update(ctx context.Context, id uint, data *authmodel.User) (error)
}
type JwtProvider interface {
	GenerateAccessToken(userID uint, email string) (string, error)
	ValidateAccessToken(tokenString string) (*jwt.Claims, error)
	GenerateRandomString(n int) (string, error)
}
type loginBiz struct{
	storage LoginStorage
	jwtProvider JwtProvider
	hasher Hasher
}
func NewLoginBusiness(storage LoginStorage, jwtProvider JwtProvider,hasher Hasher) *loginBiz{
	return &loginBiz{
		storage: storage,
		jwtProvider: jwtProvider,
		hasher: hasher,
	}
}
func (biz *loginBiz) Login(ctx context.Context,data *authmodel.UserLogin) (*authmodel.TokenResponse,error){
	user, err := biz.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	log.Println(user)
	if err != nil {
		return nil, err
	}
	if  !biz.hasher.CheckPasswordAndHash(data.Password,user.Password){
		return nil,err
	}
	
	accessToken,err := biz.jwtProvider.GenerateAccessToken(user.Id,user.Email)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	refreshToken,err:=biz.jwtProvider.GenerateRandomString(32)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	
	expiresAt :=time.Now().AddDate(0,3,0)
	user.ExpiresAt=&expiresAt
	user.RefreshToken=&refreshToken
	err=biz.storage.Update(ctx,user.Id,user)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return &authmodel.TokenResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}