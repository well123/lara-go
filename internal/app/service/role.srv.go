package service

import (
	"context"
	"github.com/google/wire"
	"goApi/internal/app/dao/role"
	"goApi/internal/app/schema"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleSrv), "*"))

type RoleSrv struct {
	RoleRepo *role.Repo
}

func (s *RoleSrv) GetCasbinRolePermissions(ctx context.Context) []schema.CasbinRoleParam {
	var roles []role.Role
	var schemaRoles []schema.CasbinRoleParam
	s.RoleRepo.GetRolePermissions(ctx, &roles)
	for _, r := range roles {
		for _, permission := range r.RolePermission {
			schemaRoles = append(schemaRoles, schema.CasbinRoleParam{
				RoleID: r.ID,
				Path:   permission.Permission.Path,
			})
		}
	}
	return schemaRoles
}
