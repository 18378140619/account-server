package api

import (
	"account-server/service"
	"account-server/service/dto"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

func (u UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError()
	if err != nil {
		return
	}
	iUser, token, err := u.Service.Login(iUserLoginDTO)
	if err != nil {
		u.Fail(err.Error())
		return
	}
	u.OK(gin.H{
		"token": token,
		"user":  iUser,
	})
}

func (u UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO}).GetError()
	if err != nil {
		return
	}
	err = u.Service.AddUser(iUserAddDTO)
	if err != nil {
		u.Fail(err.Error())
		return
	}
	u.OK(iUserAddDTO)
}

func (u UserApi) GetUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindParamsFromUri: true}).GetError()
	if err != nil {
		return
	}
	iUser, err := u.Service.GetUserById(iCommonIDDTO)
	if err != nil {
		u.Fail(err.Error())
		return
	}
	u.OK(iUser)
}

func (u UserApi) GetUserList(c *gin.Context) {
	var iUserListDTO dto.UserListDTO
	err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserListDTO}).GetError()
	if err != nil {
		return
	}
	userList, total, err := u.Service.GetUserList(iUserListDTO)
	if err != nil {
		u.Fail(err.Error())
		return
	}
	u.OKl(userList, total)
}

func (u UserApi) EditUser(c *gin.Context) {
	var iUserEditDTO dto.UserEditDTO
	err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserEditDTO}).GetError()
	if err != nil {
		return
	}
	err = u.Service.EditUser(&iUserEditDTO)
	if err != nil {
		u.Fail(err.Error())
		return
	}
	u.OK(nil)
}

func (u UserApi) DelUser(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindParamsFromUri: true}).GetError()
	if err != nil {
		return
	}
	err = u.Service.DelUser(iCommonIDDTO)
	if err != nil {
		u.Fail(err.Error())
		return
	}
	u.OK(nil)
}
