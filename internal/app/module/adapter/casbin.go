package adapter

import (
	"context"
	"fmt"
	"goApi/internal/app/service"

	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/google/wire"

	"goApi/pkg/logger"
)

var _ persist.Adapter = (*CasbinAdapter)(nil)

var CasbinAdapterSet = wire.NewSet(wire.Struct(new(CasbinAdapter), "*"), wire.Bind(new(persist.Adapter), new(*CasbinAdapter)))

type CasbinAdapter struct {
	RoleSrv *service.RoleSrv
	UserSrv *service.UserSrv
}

// Loads all policy rules from the storage.
func (a *CasbinAdapter) LoadPolicy(model casbinModel.Model) error {
	ctx := context.Background()
	err := a.loadRolePolicy(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("Load casbin role policy error: %s", err.Error())
		return err
	}

	err = a.loadUserPolicy(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("Load casbin user policy error: %s", err.Error())
		return err
	}

	return nil
}

// Load role policy (p,role_id,path,method)
func (a *CasbinAdapter) loadRolePolicy(ctx context.Context, m casbinModel.Model) error {
	rolePermissions := a.RoleSrv.GetCasbinRolePermissions(ctx)
	for _, rolePermission := range rolePermissions {
		if rolePermission.Path == "" {
			continue
		}
		_ = persist.LoadPolicyLine(fmt.Sprintf("p,%d,%s,POST", rolePermission.RoleID, rolePermission.Path), m)
	}

	return nil
}

// Load user policy (g,user_id,role_id)
func (a *CasbinAdapter) loadUserPolicy(ctx context.Context, m casbinModel.Model) error {
	userRoles := a.UserSrv.GetCasbinUserRoles(ctx)
	for _, item := range userRoles {
		_ = persist.LoadPolicyLine(fmt.Sprintf("g,%d,%d", item.UserID, item.RoleID), m)
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *CasbinAdapter) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
