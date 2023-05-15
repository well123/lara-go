//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"goApi/internal/app/api"
	"goApi/internal/app/dao"
	"goApi/internal/app/module/adapter"
	"goApi/internal/app/router"
	"goApi/internal/app/service"
	"goApi/internal/app/setup"
)

func BuildInjector() (*setup.Injector, func(), error) {
	wire.Build(
		setup.Gorm,
		dao.RepoSet,
		adapter.CasbinAdapterSet,
		api.Set,
		service.ServiceSet,
		setup.Casbin,
		setup.Cron,
		setup.Auth,
		setup.GinEngine,
		setup.InjectorSet,
		router.RouterSet,
	)
	return new(setup.Injector), nil, nil
}
