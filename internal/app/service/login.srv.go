package service

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"goApi/internal/app/dao/user"
	"goApi/internal/app/schema"
	"goApi/pkg/auth"
	"goApi/pkg/errors"
)

var LoginSet = wire.NewSet(wire.Struct(new(LoginSrv), "*"))

type LoginSrv struct {
	Auth     auth.Author
	UserRepo user.UserRepo
}

func (s *LoginSrv) GetUserByApiInfo(ctx context.Context, param *schema.LoginUserParam) (*schema.User, error) {
	return s.UserRepo.GetByApi(ctx, param)
}

func (s *LoginSrv) GenerateToken(ctx context.Context, user *schema.User) (*schema.LoginTypeInfo, error) {
	token, err := s.Auth.GenerateToken(ctx, s.formatUserInfoForToken(user))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &schema.LoginTypeInfo{
		AccessToken: token.GetAccessToken(),
		ExpiresAt:   token.GetExpiresAt(),
		AccessType:  token.GetTokenType(),
	}, nil
}

func (s *LoginSrv) formatUserInfoForToken(user *schema.User) string {
	return fmt.Sprintf("%d-%s", user.ID, user.UserName)
}
