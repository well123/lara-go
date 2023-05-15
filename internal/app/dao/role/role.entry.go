package role

import (
	"context"
	"goApi/internal/app/dao/util"
	"gorm.io/gorm"
)

type Role struct {
	util.Model
	Name           string `gorm:"size:64;default:0;comment:'角色'"`
	IsDelete       int    `gorm:"index;default:0"`
	RolePermission []RolePermission
}

func GetRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Role))
}
