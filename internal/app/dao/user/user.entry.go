package user

import (
	"context"
	"goApi/internal/app/dao/role"
	"goApi/internal/app/dao/util"
	"gorm.io/gorm"
)

func GetUserDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, db, User{})
}

type User struct {
	util.Model
	UserName  string  `gorm:"size:64;uniqueIndex;default:'';not null;"`
	RealName  string  `gorm:"size:64;index;default:'';"` // 真实姓名
	Password  string  `gorm:"size:40;default:'';"`       // 密码
	Email     *string `gorm:"size:255;"`                 // 邮箱
	Phone     *string `gorm:"size:20;"`                  // 手机号
	Status    int     `gorm:"index;default:0;"`          // 状态(1:启用 2:停用)
	ApiId     string  `gorm:"size:64;default:'';"`
	ApiSecret string  `gorm:"size:64;default:'';"`
	RoleID    int
	Role      role.Role
}
