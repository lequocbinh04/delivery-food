package component

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

type AppCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *AppCtx {
	return &AppCtx{db: db}
}

func (ctx *AppCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

//func (ctx *AppCtx) SecretKey() string { return ctx.secretKey }
