//go:build wireinject
// +build wireinject

package di

import (
	"app/api/delete-instance/infra/data"
	"app/api/delete-instance/infra/web"
	"app/api/delete-instance/usecase"

	"github.com/google/wire"
)

var DeleteInstance = wire.NewSet(
	web.NewHandler,
	data.NewServiceRepository,
	usecase.NewService,
)
