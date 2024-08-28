package authtransport

import (
	"go-base/component/appctx"
	"go-base/component/hasher"
	authbiz "go-base/module/auth/biz"
	authmodel "go-base/module/auth/model"
	authstorage "go-base/module/auth/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainConnection()
		var data authmodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := authstorage.NewSQLStore(db)
		hasher:=hasher.NewHasher()
		biz := authbiz.NewRegisterBusiness(store,hasher)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK,gin.H{
			"message":"user register success",
			"isSuccess":true,
		})
	}
}