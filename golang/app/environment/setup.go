package environment

import (
	"app/di"
	"app/environment/waf"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Application はEchoアプリケーションを変数に持つ構造体です。
type Application struct {
	App *echo.Echo
}
type FileLogger struct {
	file *os.File
}

// NewApp はEchoアプリケーションを作成します。
func NewApp(db *gorm.DB) *Application {
	echo := waf.NewEcho()

	handler := di.InitializeHandler(db, echo)

	waf.NewRouter(echo, handler)

	return &Application{
		App: echo,
	}

}

// Start は与えられたportでEchoアプリケーションをスタートします。
func (a *Application) Start(port string) {
	// サーバ起動
	a.App.Logger.Fatal(a.App.Start(":" + port))
}
