package router

import (
	"github.com/gin-gonic/gin"
	"goApi/internal/app/config"
	"goApi/internal/app/middleware"
	"net/http"
)

func (a *Router) GetWebRoutes(g *gin.RouterGroup) {
	if config.C.JWTAuth.Enable {
		g.Use(middleware.AuthMiddleware(a.Auth, middleware.AllowPathPrefixSkipper("/v1/login")))
	}
	v1 := g.Group("/v1")
	{
		v1.POST("/login", func(context *gin.Context) {
			context.String(http.StatusOK, "Hello, Gin!")
		})
		v1.POST("/logout", func(context *gin.Context) {
			context.String(http.StatusOK, "Hello, Gin!")
		})
	}
}
