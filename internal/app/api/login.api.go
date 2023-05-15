package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"goApi/internal/app/schema"
	"goApi/internal/app/service"
	"goApi/internal/app/util"
	"goApi/pkg/logger"
)

var LoginSet = wire.NewSet(wire.Struct(new(LoginApi), "*"))

type LoginApi struct {
	LoginSrv *service.LoginSrv
}

func (a *LoginApi) GenerateToken(c *gin.Context) {
	ctx := c.Request.Context()
	var loginUserParam schema.LoginUserParam
	if err := util.ParseJson(c, &loginUserParam); err != nil {
		util.ResError(c, err)
		return
	}
	info, err := a.LoginSrv.GetUserByApiInfo(ctx, &loginUserParam)
	if err != nil {
		util.ResError(c, err)
		return
	}
	token, err := a.LoginSrv.GenerateToken(ctx, info)
	if err != nil {
		util.ResError(c, err)
		return
	}

	ctx = logger.NewTagCtx(ctx, "__login__")
	ctx = logger.NewUserIDCtx(ctx, info.ID)
	ctx = logger.NewUserNameCtx(ctx, info.UserName)
	logger.WithContext(ctx).Infof("login")

	util.ResSuccess(c, token)
}
