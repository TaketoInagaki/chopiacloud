package web

import (
	"app/api/create-instance/domain"
	"app/api/create-instance/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	s usecase.Service
}

// NewHandler は初期化済みのHandlerを返す。
func NewHandler(s usecase.Service) Handler {
	return &handler{s: s}
}

// Handler はHandlerのインターフェース
type Handler interface {
	CreateInstance(c echo.Context) error
}

func (h *handler) CreateInstance(c echo.Context) error {

	// リクエストパラメータ取得
	param := domain.RequestParam{}

	// 引数の値のフォーマットチェック
	if err := c.Bind(&param); err != nil {
		return err
	}
	if param.SshKeyId == 0 {
		return echo.NewHTTPError(400)
	}

	// インスタンス作成
	response, err := h.s.CreateInstance(param.SshKeyId)
	if err != nil {
		if _, ok := err.(*domain.NotFoundError); ok {
			return echo.NewHTTPError(404)
		}
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}