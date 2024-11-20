//go:build wireinject
// +build wireinject

package di

import (
	"app/api/upload-key/infra/data"
	"app/api/upload-key/infra/web"
	"app/api/upload-key/usecase"

	"github.com/google/wire"
)

var UploadKey = wire.NewSet(
	web.NewHandler,
	data.NewServiceRepository,
	usecase.NewService,
)
