//go:build wireinject
// +build wireinject

package di

import (
	"app/api/start-instance/infra/data"
	"app/api/start-instance/infra/web"
	"app/api/start-instance/usecase"

	"github.com/google/wire"
)

var StartInstance = wire.NewSet(
	web.NewHandler,
	data.NewServiceRepository,
	usecase.NewService,
)
