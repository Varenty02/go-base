package authtransport

import (
	common "go-base/commons"
	"go-base/component/appctx"
	"go-base/component/jwt"
	authbiz "go-base/module/auth/biz"
	authmodel "go-base/module/auth/model"
	authstorage "go-base/module/auth/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateToken(ctx appctx.AppContext) gin.HandlerFunc{
	return gin.HandlerFunc(func(c *gin.Context) {
		var tokenPairRequest=authmodel.TokenRequest{}
		if err:=c.ShouldBind(&tokenPairRequest);err!=nil{
			panic(common.ErrInvalidRequest(err))
		}
		db:=authstorage.NewSQLStore(ctx.GetMainConnection())
		jwtProvider:=jwt.NewJWTProvider(ctx.GetSecretKey())
		biz:=authbiz.NewGenerateTokenBiz(jwtProvider,db)
		tokenPair, err := biz.GenerateToken(c.Request.Context(),&tokenPairRequest)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK,tokenPair)
	})

}