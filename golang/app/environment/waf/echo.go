package waf

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// echoインスタンスを新規作成
func NewEcho() *echo.Echo {
	echo := echo.New()

	// APIアクセスIDをcontext内で使用できるようにする。
	echo.Use(middleware.RequestID())

	echo.Use(middleware.Recover())

	// CROS対応
	echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("CRM_URL")},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	return echo
}
