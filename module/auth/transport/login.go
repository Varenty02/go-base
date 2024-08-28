package authtransport

import (
	common "go-base/commons"
	"go-base/component/appctx"
	"go-base/component/hasher"
	"go-base/component/jwt"
	authbiz "go-base/module/auth/biz"
	authmodel "go-base/module/auth/model"
	authstorage "go-base/module/auth/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData authmodel.UserLogin
		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainConnection()
		tokenProvider := jwt.NewJWTProvider(appCtx.GetSecretKey())
		store := authstorage.NewSQLStore(db)
		appHasher := hasher.NewHasher()
		business := authbiz.NewLoginBusiness(store, tokenProvider, appHasher)
		tokenPair, err := business.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, tokenPair)

	}
}