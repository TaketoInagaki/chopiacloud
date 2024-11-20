//go:build wireinject
// +build wireinject

package di

import (
	"app/api/create-instance/infra/data"
	"app/api/create-instance/infra/web"
	"app/api/create-instance/usecase"

	"github.com/google/wire"
)

var CreateInstance = wire.NewSet(
	web.NewHandler,
	data.NewServiceRepository,
	usecase.NewService,
)
