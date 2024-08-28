package middleware

import (
	"errors"
	"go-base/component/appctx"
	"go-base/component/jwt"
	authstorage "go-base/module/auth/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("Wrong authenticate header")
	}
	return parts[1], nil
}
func RequireAuth(ctx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewJWTProvider(ctx.GetSecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		db := ctx.GetMainConnection()
		store := authstorage.NewSQLStore(db)
		claims, err := tokenProvider.ValidateAccessToken(token)
		if err != nil {
			panic(err)
		}
		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": claims.UserID})
		if err != nil {
			panic(err)
		}
		// if user.Status == 0 {
		// 	panic(commons.ErrNoPermission(errors.New("user has been deleted or banned")))
		// }
		c.Set("current_user", user)
		c.Next()
	}
}