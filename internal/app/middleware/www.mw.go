package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func WWWMiddleware(dir string, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		if SkipHandler(context, skippers...) {
			context.Next()
			return
		}
		path := context.Request.URL.Path
		path = filepath.Join(dir, filepath.FromSlash(path))
		_, err := os.Stat(path)
		if err != nil && os.IsNotExist(err) {
			path = filepath.Join(dir, "index.html")
		}
		context.File(path)
		context.Abort()
	}
}
