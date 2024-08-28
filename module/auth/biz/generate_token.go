package authbiz

import (
	"context"
	"errors"
	authmodel "go-base/module/auth/model"
	"time"
)

type RefreshTokenStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*authmodel.User, error)
	Update(ctx context.Context, id uint, data *authmodel.User) (error)
}

type generateTokenBiz struct {
	jwtProvider JwtProvider
	refreshTokenStorage RefreshTokenStorage

}
func NewGenerateTokenBiz(jwtProvider JwtProvider, refreshTokenStorage RefreshTokenStorage) *generateTokenBiz{
	return &generateTokenBiz{
		jwtProvider: jwtProvider,
		refreshTokenStorage: refreshTokenStorage,
	}
}
func (biz *generateTokenBiz) GenerateToken(ctx context.Context,tokenPair *authmodel.TokenRequest)(*authmodel.TokenResponse,error) {
	claims,err:=biz.jwtProvider.ValidateAccessToken(tokenPair.AccessToken)
	if err!=nil{
		return nil,err
	}
	user,err:=biz.refreshTokenStorage.FindUser(ctx,map[string]interface{}{"email":claims.Email})
	if err!=nil||*user.RefreshToken!=tokenPair.RefreshToken||user.ExpiresAt.Before(time.Now()){
		return nil,errors.New("Invalid refresh token")
	}
	newAccessToken,err:=biz.jwtProvider.GenerateAccessToken(user.Id,user.Email)
	if err!=nil{
		return nil ,err
	}
	newRefreshToken,err:=biz.jwtProvider.GenerateRandomString(32)
	if err!=nil{
		return nil,err
	}
	user.RefreshToken=&newRefreshToken
	expiresAt :=time.Now().AddDate(0,3,0)
	user.ExpiresAt=&expiresAt
	biz.refreshTokenStorage.Update(ctx,user.Id,user)
	return &authmodel.TokenResponse{
		AccessToken: newAccessToken,
		RefreshToken: newRefreshToken,
	},nil
}