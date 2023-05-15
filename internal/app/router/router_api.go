package router

import (
	"github.com/gin-gonic/gin"
	"goApi/internal/app/config"
	"goApi/internal/app/middleware"
	"net/http"
)

func (a *Router) GetApiRoutes(g *gin.RouterGroup) {

	if config.C.JWTAuth.Enable {
		g.Use(middleware.AuthMiddleware(a.Auth, middleware.AllowPathPrefixSkipper("/api/v1/login")))
	}

	if config.C.CasBin.Enable {
		g.Use(middleware.CasbinMiddleware(a.CasbinEnforcer, middleware.AllowPathPrefixSkipper("/api/v1/login")))
	}

	g.Use(middleware.RateLimiterMiddleware())
	prefixes := a.Prefixes()

	g.Use(middleware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))
	g.Use(middleware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	v1 := g.Group("/v1")
	{
		v1.POST("/login", a.LoginApi.GenerateToken)
		v1.POST("/logout", func(context *gin.Context) {
			context.String(http.StatusOK, "Hello, Gin!")
		})
	}

}
