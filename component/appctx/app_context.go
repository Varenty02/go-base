package appctx

import (
	"gorm.io/gorm"
)

type appContext struct {
	db        *gorm.DB
	secretKey string
}
type AppContext interface {
	GetMainConnection() *gorm.DB
	GetSecretKey() string
}

func NewAppContext(db *gorm.DB, secretKey string) *appContext {
	return &appContext{
		db:        db,
		secretKey: secretKey,
	}
}
func (appCtx *appContext) GetMainConnection() *gorm.DB { return appCtx.db }
func (appCtx *appContext) GetSecretKey() string        { return appCtx.secretKey }