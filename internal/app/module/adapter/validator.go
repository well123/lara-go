package adapter

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	validator2 "goApi/internal/app/validator"
	"goApi/pkg/logger"
	"strings"
	"sync"
)

var errorMsg map[string]func(filed, param string) string

type CustomValidator struct {
	once sync.Once
	*validator.Validate
}

func (cv *CustomValidator) RegisterValidation(tag string, fn validator.Func) error {
	return cv.Validate.RegisterValidation(tag, fn)
}

func (cv *CustomValidator) ValidateStruct(s interface{}) error {
	cv.lazyInit()
	return cv.Struct(s)
}

func (cv *CustomValidator) Engine() interface{} {
	cv.lazyInit()
	return cv.Validate
}

func (cv *CustomValidator) lazyInit() {
	cv.once.Do(func() {
		cv.Validate = validator.New()
		cv.Validate.SetTagName("validator")
	})
}

func init() {
	errorMsg = validator2.InitValidatorErrors()
}

func (cv *CustomValidator) Struct(s any) error {
	err := cv.Validate.Struct(s)
	ctx := context.Background()
	isDebug := gin.Mode() == "debug"
	if err != nil {
		var errMsgs []string
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e.Tag())
			if fn, ok := errorMsg[e.Tag()]; ok {
				errMsgs = append(errMsgs, fn(e.Field(), e.Param()))
				continue
			} else {
				if isDebug {
					panic(fmt.Sprintf("unknown validator tag: %s", e.Tag()))
				} else {
					logger.WithContext(ctx).Errorf("unknown validator tag: %s", e.Tag())
				}
			}
		}
		return fmt.Errorf(strings.Join(errMsgs, ","))
	}
	return nil
}
