package middleware

import (
	"github.com/gin-gonic/gin"
	"goApi/internal/app/util"
	"goApi/pkg/errors"
)

func NoMethodHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		util.ResError(context, errors.ErrMethodNotFound)
	}
}

func NotFoundHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		util.ResError(context, errors.ErrNotFound)
	}
}

type SkipperFunc func(*gin.Context) bool

func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(context *gin.Context) bool {
		path := context.Request.URL.Path
		pathLen := len(path)
		for _, prefix := range prefixes {
			if pl := len(prefix); pathLen >= pl && path[:pl] == prefix {
				return true
			}
		}
		return false
	}
}

func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(context *gin.Context) bool {
		path := context.Request.URL.Path
		pathLen := len(path)
		for _, prefix := range prefixes {
			if pl := len(prefix); pathLen >= pl && path[:pl] == prefix {
				return false
			}
		}
		return true
	}
}

func SkipHandler(context *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(context) {
			return true
		}
	}
	return false
}

func EmptyHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
