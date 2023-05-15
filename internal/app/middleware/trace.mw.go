package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"goApi/internal/app/util"
	"goApi/pkg/logger"
	"goApi/pkg/trace"
)

func TraceMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if SkipHandler(ctx, skippers...) {
			ctx.Next()
			return
		}

		traceID := ctx.GetHeader("X-Request-ID")
		if traceID == "" {
			traceID = trace.NewTraceID()
		}

		c := context.WithValue(ctx, &util.TraceIDCtx{}, traceID)
		c = logger.NewTraceIdCtx(c, traceID)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Writer.Header().Set("X-Trace-ID", traceID)
		ctx.Next()
	}
}
