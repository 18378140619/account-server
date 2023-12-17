package api

import (
	"account-server/service"
	"account-server/service/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginApi struct {
	BaseApi
	Service *service.LoginService
}

func NewLoginApi() LoginApi {
	return LoginApi{
		BaseApi: NewBaseApi(),
		Service: service.NewLoginService(),
	}
}

func (l LoginApi) WxCodeGetOpenID(c *gin.Context) {
	var iUserCodeDTO dto.UserCodeDTO
	err := l.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserCodeDTO}).GetError()
	if err != nil {
		return
	}
	var s string
	s += "https://api.weixin.qq.com/sns/jscode2session?appid="
	s += "wx64c12a1645d44164"
	s += "&secret="
	s += "f87abde966e34da5c906dfeaef0ba859"
	s += "&js_code="
	s += iUserCodeDTO.Code
	s += "&grant_type=authorization_code"
	fmt.Printf("%+v", s)
	response, err := http.Get(s)
	if err != nil {
		l.Fail(err.Error())
	}
	fmt.Printf("%+v", response)
	l.OK(response)
}
