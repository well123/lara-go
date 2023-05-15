package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goApi/pkg/auth/jwt"
	"testing"
)

func TestRouter(t *testing.T) {
	a := &Router{Auth: &jwt.JWTAuth{}}
	c := gin.New()
	a.registerApi(c)
	println("cc")
	for _, route := range c.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}
}
