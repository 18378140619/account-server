package api

import (
	"github.com/gin-gonic/gin"
)

type ResponseResult struct {
	Msg   string      `json:"msg,omitempty"`  // 错误描述
	Code  int         `json:"code,omitempty"` // 错误码
	Data  interface{} `json:"data,omitempty"` // 返回数据
	Total int64       `json:"total,omitempty"`
}

type ResponseErrResult struct {
	Msg  string      `json:"msg,omitempty"`  // 错误描述
	Code int         `json:"code,omitempty"` // 错误码
	Data interface{} `json:"data,omitempty"` // 返回数据
}

// ResponseSuccess 成功响应，不带msg形参，默认success
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(SUCCESS, &ResponseResult{
		Msg:  "success",
		Code: SUCCESS,
		Data: data,
	})
}

// BusinessErrApiResult 普通业务异常
func BusinessErrApiResult(c *gin.Context, msg string) {
	c.JSON(SUCCESS, &ResponseResult{
		Msg:  msg,
		Code: ERROR,
		Data: nil,
	})
}

func ErrApiResult(c *gin.Context, msg string, code int, data interface{}) {
	c.JSON(SUCCESS, &ResponseResult{
		Msg:  msg,
		Code: code,
		Data: data,
	})
}

func SuccessApiResult(c *gin.Context, msg string, code int, data interface{}, total int64) {
	c.JSON(SUCCESS, &ResponseResult{
		Msg:   msg,
		Code:  code,
		Data:  data,
		Total: total,
	})
}
