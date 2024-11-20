//go:build wireinject
// +build wireinject

package di

import (
	"app/api/stop-instance/infra/data"
	"app/api/stop-instance/infra/web"
	"app/api/stop-instance/usecase"

	"github.com/google/wire"
)

var StopInstance = wire.NewSet(
	web.NewHandler,
	data.NewServiceRepository,
	usecase.NewService,
)
