//go:build wireinject
// +build wireinject

package di

import (
	"app/environment/waf"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitializeHandler(d *gorm.DB, e *echo.Echo) waf.Handler {
	wire.Build(
		waf.NewHandler,
		CreateInstance,
		StartInstance,
		StopInstance,
		DeleteInstance,
		UploadKey,
	)

	return waf.Handler{}
}
