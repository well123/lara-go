// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"goApi/internal/app/api"
	"goApi/internal/app/dao/role"
	"goApi/internal/app/dao/user"
	"goApi/internal/app/module/adapter"
	"goApi/internal/app/router"
	"goApi/internal/app/service"
	"goApi/internal/app/setup"
)

// Injectors from wire.go:

func BuildInjector() (*setup.Injector, func(), error) {
	author, cleanup, err := setup.Auth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := setup.Gorm()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	repo := &role.Repo{
		DB: db,
	}
	roleSrv := &service.RoleSrv{
		RoleRepo: repo,
	}
	userRepo := &user.UserRepo{
		DB: db,
	}
	userSrv := &service.UserSrv{
		UserRepo: userRepo,
	}
	casbinAdapter := &adapter.CasbinAdapter{
		RoleSrv: roleSrv,
		UserSrv: userSrv,
	}
	syncedEnforcer, cleanup3, err := setup.Casbin(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userApi := &api.UserApi{
		UserSrv: userSrv,
	}
	userUserRepo := user.UserRepo{
		DB: db,
	}
	loginSrv := &service.LoginSrv{
		Auth:     author,
		UserRepo: userUserRepo,
	}
	loginApi := &api.LoginApi{
		LoginSrv: loginSrv,
	}
	routerRouter := &router.Router{
		Auth:           author,
		CasbinEnforcer: syncedEnforcer,
		UserApi:        userApi,
		LoginApi:       loginApi,
	}
	engine := setup.GinEngine(routerRouter)
	cron, cleanup4, err := setup.Cron()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	injector := &setup.Injector{
		Engine:         engine,
		Auth:           author,
		DB:             db,
		CasbinEnforcer: syncedEnforcer,
		Cron:           cron,
	}
	return injector, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
