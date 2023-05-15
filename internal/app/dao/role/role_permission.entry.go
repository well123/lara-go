package role

import (
	"context"
	"goApi/internal/app/dao/permission"
	"goApi/internal/app/dao/util"
	"gorm.io/gorm"
)

type RolePermission struct {
	util.Model
	RoleID       int `gorm:"index;default:0"`
	PermissionID int `gorm:"index;default:0"`
	Permission   permission.Permission
}

func GetRolePermissionDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(RolePermission))
}
