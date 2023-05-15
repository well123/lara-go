package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goApi/internal/app/schema"
	"goApi/pkg/errors"
	"goApi/pkg/logger"
	"net/http"
	"strings"
)

var (
	prefix     = "go-api"
	ResBodyKey = prefix + "/res-body"
	ReqBodyKey = prefix + "/req-body"
)

func ParseJson(c *gin.Context, obj any) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Parse request json failed: %s", err.Error()))
	}
	return nil
}

// Parse query parameter to struct
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Parse request query failed: %s", err.Error()))
	}
	return nil
}

// Parse body form data to struct
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Parse request form failed: %s", err.Error()))
	}
	return nil
}

func GetToken(c *gin.Context) string {
	authorization := c.GetHeader("Authorization")
	prefix := "Bearer "
	if authorization != "" && strings.HasPrefix(authorization, prefix) {
		return authorization[len(prefix):]
	}
	return ""
}

func ResJson(c *gin.Context, status int, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, data)
	c.Data(status, "application/json; charset=utf-8", data)
	c.Abort()
}

func ResSuccess(c *gin.Context, v any) {
	ResJson(c, http.StatusOK, v)
}

func ResError(c *gin.Context, err error, status ...int) {
	ctx := c.Request.Context()
	var res *errors.ResponseError

	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.ErrInternalServer)
			res.ERR = err
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.Status = status[0]
	}

	if err := res.ERR; err != nil {
		if res.Message == "" {
			res.Message = err.Error()
		}

		if status := res.Status; status >= 400 && status < 500 {
			logger.WithContext(ctx).Warnf(err.Error())
		} else if status >= 500 {
			logger.WithContext(logger.WithStackContext(ctx, err)).Errorf(err.Error())
		}
	}

	errorItem := schema.ErrorItem{
		Code:    res.Code,
		Message: res.Message,
	}
	ResJson(c, res.Status, schema.ErrorResult{Error: errorItem, Status: false})
}
