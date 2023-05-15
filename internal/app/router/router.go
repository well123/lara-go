package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"goApi/internal/app/api"
	"goApi/pkg/auth"
	"os"
	"reflect"
	"strings"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type Router struct {
	Auth           auth.Author
	CasbinEnforcer *casbin.SyncedEnforcer
	UserApi        *api.UserApi
	LoginApi       *api.LoginApi
}

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

func (a *Router) Register(app *gin.Engine) error {
	a.registerApi(app)
	return nil
}

func (a *Router) registerApi(app *gin.Engine) {

	files := getWdFiles()
	for _, file := range files {
		g := app.Group("/" + file)
		a.addGroup(g, snakeToHump(file))
	}
}
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

func getWdFiles() []string {
	files, _ := os.ReadDir("./internal/app/router")

	var temp []string
	for _, file := range files {
		fileName := file.Name()
		if !file.IsDir() && strings.HasSuffix(fileName, ".go") && strings.HasPrefix(fileName, "router") && !strings.HasSuffix(fileName, "_test.go") && fileName != "router.go" {
			temp = append(temp, fileName[7:len(fileName)-3])
		}
	}
	return temp
}

func (a *Router) addGroup(g *gin.RouterGroup, file string) {
	funcName := "Get" + file + "Routes"
	method := reflect.ValueOf(a).MethodByName(funcName)

	if method.Kind() != reflect.Func {
		return
	}
	method.Call([]reflect.Value{
		reflect.ValueOf(g),
	})
}

// 将蛇形字符串转换成驼峰字符串
func snakeToHump(s string) string {
	strs := strings.Split(s, "_")
	for index, str := range strs {
		strs[index] = strings.ToUpper(str[:1]) + str[1:]
	}
	return strings.Join(strs, "")
}
