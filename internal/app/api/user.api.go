package api

import (
	"github.com/google/wire"
	"goApi/internal/app/service"
)

var UserSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserSrv *service.UserSrv
}
