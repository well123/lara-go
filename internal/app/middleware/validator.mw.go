package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goApi/internal/app/module/adapter"
)

func ValidatorMiddleware() gin.HandlerFunc {
	binding.Validator = &adapter.CustomValidator{}

	return func(context *gin.Context) {
		context.Next()
	}
}
