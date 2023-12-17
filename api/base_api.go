package api

import (
	"account-server/global"
	"account-server/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromUri bool
}

func (b *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error

	//绑定请求上下文
	b.Ctx = option.Ctx

	//绑定请求数据
	if option.DTO != nil {
		if option.BindParamsFromUri {
			errResult = b.Ctx.ShouldBindUri(option.DTO)
		} else {
			errResult = b.Ctx.ShouldBind(option.DTO)
		}
	}
	if errResult != nil {
		errResult = b.ParseValidateErrors(errResult, option.DTO)
		b.addError(errResult)
		b.Fail(b.GetError().Error())
	}
	return b
}

// ParseValidateErrors 错误处理 转中文
func (b *BaseApi) ParseValidateErrors(errs error, target any) error {
	var errResult error

	var errValidation validator.ValidationErrors
	ok := errors.As(errs, &errValidation)
	if !ok {
		return errs
	}

	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s：%s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}

func (b *BaseApi) addError(errNew error) {
	b.Errors = utils.AppendError(b.Errors, errNew)
}

func (b *BaseApi) GetError() error {
	return b.Errors
}

func (b *BaseApi) OK(data interface{}) {
	ResponseSuccess(b.Ctx, data)
}

func (b *BaseApi) Fail(msg string) {
	BusinessErrApiResult(b.Ctx, msg)
}

func (b *BaseApi) Fails(msg string, code int, data interface{}) {
	if code == 0 {
		code = 500
	}
	ErrApiResult(b.Ctx, msg, code, data)
}

func (b *BaseApi) OKl(data interface{}, total int64) {
	SuccessApiResult(b.Ctx, "操作成功", SUCCESS, data, total)
}
