package setup

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"goApi/internal/app/cron"
	"goApi/pkg/auth"
	"gorm.io/gorm"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine         *gin.Engine
	Auth           auth.Author
	DB             *gorm.DB
	CasbinEnforcer *casbin.SyncedEnforcer
	Cron           *cron.Cron
}
