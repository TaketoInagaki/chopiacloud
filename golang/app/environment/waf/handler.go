package waf

import (
	CreateInstance "app/api/create-instance/infra/web"
	DeleteInstance "app/api/delete-instance/infra/web"
	StartInstance "app/api/start-instance/infra/web"
	StopInstance "app/api/stop-instance/infra/web"
	UploadKey "app/api/upload-key/infra/web"
)

// Handler の構造体。
// APIを追加する場合はフィールドを追加する。
type Handler struct {
	CreateInstance CreateInstance.Handler
	StartInstance  StartInstance.Handler
	StopInstance   StopInstance.Handler
	DeleteInstance DeleteInstance.Handler
	UploadKey      UploadKey.Handler
}

// NewHandler Handlerを初期化して返す。
// APIを追加する場合は引数と構造体に値をセットする行を追加する。
func NewHandler(
	CreateInstance CreateInstance.Handler,
	StartInstance StartInstance.Handler,
	StopInstance StopInstance.Handler,
	DeleteInstance DeleteInstance.Handler,
	UploadKey UploadKey.Handler,

) Handler {
	return Handler{
		CreateInstance: CreateInstance,
		StartInstance:  StartInstance,
		StopInstance:   StopInstance,
		DeleteInstance: DeleteInstance,
		UploadKey:      UploadKey,
	}
}
