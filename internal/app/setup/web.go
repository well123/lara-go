package setup

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goApi/internal/app/config"
	"goApi/internal/app/middleware"
	"goApi/internal/app/router"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) ValidateStruct(obj interface{}) error {
	if err := cv.validator.Struct(obj); err != nil {
		return err
	}
	return nil
}

func (cv *CustomValidator) Engine() interface{} {
	return cv.validator
}

func GinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)
	app := gin.New()

	app.NoRoute(middleware.NotFoundHandler())
	app.NoMethod(middleware.NoMethodHandler())

	prefixes := r.Prefixes()
	app.Use(middleware.ValidatorMiddleware())
	app.Use(middleware.RecoverMiddleware())

	// CORS
	if config.C.CORS.Enable {
		app.Use(middleware.CorsMiddleware())
	}

	// GZIP
	if config.C.GZIP.Enable {
		app.Use(gzip.Gzip(gzip.BestCompression,
			gzip.WithExcludedExtensions(config.C.GZIP.ExcludedExtensions),
			gzip.WithExcludedPaths(config.C.GZIP.ExcludedPaths),
		))
	}

	_ = r.Register(app)

	// Swagger
	if config.C.Swagger {
		//目前不打算接入
		//app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	if dir := config.C.WWW; dir != "" {
		app.Use(middleware.WWWMiddleware(config.C.WWW, middleware.AllowPathPrefixSkipper(prefixes...)))
	}
	return app
}
