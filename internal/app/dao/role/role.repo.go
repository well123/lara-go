package role

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RoleSet = wire.NewSet(wire.Struct(new(Repo), "*"))

type Repo struct {
	DB *gorm.DB
}

func (a *Repo) GetRolePermissions(ctx context.Context, roles *[]Role) {
	GetRoleDB(ctx, a.DB).Preload("RolePermission.Permission").Where("is_delete=?", 0).Find(&roles)
}
