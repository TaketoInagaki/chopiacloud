package web

import (
	"app/api/upload-key/domain"
	"app/api/upload-key/usecase"
	"net/http"
	"strconv"

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
	UploadKey(c echo.Context) error
}

func (h *handler) UploadKey(c echo.Context) error {

	// リクエストパラメータ取得
	param := domain.RequestParam{}

	// 引数の値のフォーマットチェック
	if err := c.Bind(&param); err != nil {
		return err
	}
	if param.PublicKey == "" || param.Name == "" {
		return echo.NewHTTPError(400)
	}

	// インスタンス作成
	id, err := h.s.UploadKey(param)
	if err != nil {
		if _, ok := err.(*domain.NotFoundError); ok {
			return echo.NewHTTPError(404)
		}
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"id": strconv.FormatUint(uint64(id), 10)})
}
