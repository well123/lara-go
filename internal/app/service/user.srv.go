package service

import (
	"context"
	"github.com/google/wire"
	"goApi/internal/app/dao/user"
	"goApi/internal/app/schema"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSrv), "*"))

type UserSrv struct {
	UserRepo *user.UserRepo
}

func (s *UserSrv) GetCasbinUserRoles(ctx context.Context) []schema.CasbinUserRoleParam {
	var users []user.User
	var schemaUsers []schema.CasbinUserRoleParam
	s.UserRepo.GetUserRole(ctx, &users)
	for _, u := range users {
		schemaUsers = append(schemaUsers, schema.CasbinUserRoleParam{
			RoleID: u.RoleID,
			UserID: u.ID,
		})
	}
	return schemaUsers
}
