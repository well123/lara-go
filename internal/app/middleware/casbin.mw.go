package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"goApi/internal/app/config"
	"goApi/internal/app/util"
	"goApi/pkg/errors"
	"strconv"
)

func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := config.C.CasBin
	if !cfg.Enable {
		return EmptyHandler()
	}

	return func(context *gin.Context) {

		if SkipHandler(context, skippers...) {
			context.Next()
			return
		}

		p := context.Request.URL.Path
		m := context.Request.Method
		userID := util.FromUint64Key(context.Request.Context(), util.UserIDCtx{})
		enforce, err := enforcer.Enforce(strconv.FormatUint(userID, 10), p, m)
		if err != nil {
			util.ResError(context, errors.WithStack(err))
			return
		} else if !enforce {
			util.ResError(context, errors.ErrNoPerm)
			return
		}
		context.Next()
	}
}
