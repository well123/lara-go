package user

import (
	"context"
	"github.com/google/wire"
	"goApi/internal/app/dao/util"
	"goApi/internal/app/schema"
	util2 "goApi/internal/app/util"
	"goApi/pkg/util/hash"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserRepo), "*"))

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) GetByApi(ctx context.Context, queries *schema.LoginUserParam) (*schema.User, error) {
	var user User
	db := GetUserDB(ctx, r.DB)
	db.Where("api_id=?", queries.ApiId)
	db.Where("api_secret=?", hash.SHA1String(queries.ApiSecret))
	first, err := util.First(db, &user)
	if err != nil {
		return nil, err
	}
	if !first {
		return nil, nil
	}
	schemaUser := new(schema.User)
	return schemaUser, util2.Copy(user, schemaUser)
}

func (r *UserRepo) GetUserRole(ctx context.Context, users *[]User) {
	GetUserDB(ctx, r.DB).Preload("Role").Find(users)
}
