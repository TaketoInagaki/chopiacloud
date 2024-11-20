package waf

import (
	"github.com/labstack/echo/v4"
)

// NewRouterはEchoのルーティングの設定
func NewRouter(e *echo.Echo, h Handler) {
	// インスタンス
	e.POST("/create-instance", h.CreateInstance.CreateInstance)
	e.POST("/start-instance", h.StartInstance.StartInstance)
	e.POST("/stop-instance", h.StopInstance.StopInstance)
	e.DELETE("/delete-instance", h.DeleteInstance.DeleteInstance)

	// SSH
	e.POST("/upload-key", h.UploadKey.UploadKey)
}
