package permission

import (
	"context"
	"goApi/internal/app/dao/util"
	"gorm.io/gorm"
)

type Permission struct {
	util.Model
	Name      string `gorm:"size:64;default:'';not null;comment:'权限名称';"`
	Code      string `gorm:"size:128;default:'';not null;comment:'编码';"`
	Path      string `gorm:"size:255;default:'';comment:'路由';"`
	Icon      string `gorm:"size:64;default:'';comment:'图标';"`
	IsShow    int    `gorm:"index;default:0;comment:'是否显示'"`
	ParentId  int    `gorm:"index;default:0;comment:'父级权限'"`
	Component string `gorm:"size:64;default:0;comment:'组件'"`
	Sort      int    `gorm:"index;default:0"`
	IsDelete  int    `gorm:"index;default:0"`
}

func GetPermissionDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Permission))
}
