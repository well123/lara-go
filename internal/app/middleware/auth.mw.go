package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"goApi/internal/app/util"
	"goApi/pkg/auth"
	"goApi/pkg/errors"
	"goApi/pkg/logger"
	"strconv"
	"strings"
)

type userInfo struct {
	userID   uint64
	userName string
}

func warpUserAuthContext(ctx *gin.Context, userInfo userInfo) {
	c := context.WithValue(ctx.Request.Context(), util.UserNameCtx{}, userInfo.userName)
	c = context.WithValue(c, util.UserIDCtx{}, userInfo.userID)
	logger.NewUserIDCtx(c, userInfo.userID)
	logger.NewUserNameCtx(c, userInfo.userName)
	ctx.Request = ctx.Request.WithContext(c)
}

func AuthMiddleware(a auth.Author, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		if SkipHandler(context, skippers...) {
			context.Next()
			return
		}
		userSubject, err := a.ParseUserID(context, util.GetToken(context))
		if err != nil {
			if err != auth.ErrInvalidToken {
				util.ResError(context, errors.WithStack(err))
				return
			}
			util.ResError(context, err)
			return
		}
		index := strings.Index(userSubject, "-")

		if index < 0 {
			util.ResError(context, errors.ErrInternalToken)
			return
		}
		userID, _ := strconv.ParseUint(userSubject[:index], 10, 64)
		warpUserAuthContext(context, userInfo{
			userID:   userID,
			userName: userSubject[index+1:],
		})
		context.Next()
	}
}
