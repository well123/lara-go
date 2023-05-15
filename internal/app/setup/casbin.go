package setup

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"goApi/internal/app/config"
	"time"
)

func Casbin(adapter persist.Adapter) (*casbin.SyncedEnforcer, func(), error) {
	cfg := config.C.CasBin
	if cfg.Model == "" {
		return new(casbin.SyncedEnforcer), nil, nil
	}

	m, err := casbin.NewSyncedEnforcer(cfg.Model)
	if err != nil {
		return nil, nil, err
	}
	m.EnableLog(cfg.Debug)

	err = m.InitWithModelAndAdapter(m.GetModel(), adapter)
	if err != nil {
		return nil, nil, err
	}
	m.EnableEnforce(cfg.Enable)

	var clearFunc func()
	if cfg.AutoLoad {
		m.StartAutoLoadPolicy(time.Duration(cfg.AutoLoadInternal) * time.Second)
		clearFunc = func() {
			m.StopAutoLoadPolicy()
		}
	}
	return m, clearFunc, nil
}
