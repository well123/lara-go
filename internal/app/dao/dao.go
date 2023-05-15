package dao

import (
	"github.com/google/wire"
	"goApi/internal/app/config"
	"goApi/internal/app/dao/permission"
	"goApi/internal/app/dao/role"
	"goApi/internal/app/dao/user"
	"gorm.io/gorm"
	"strings"
)

var RepoSet = wire.NewSet(
	user.UserSet,
	role.RoleSet,
)

func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	return db.AutoMigrate(
		new(user.User),
		new(role.Role),
		new(role.RolePermission),
		new(permission.Permission),
	)
}
