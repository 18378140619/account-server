package service

import (
	"account-server/dao"
	"account-server/global"
	"account-server/global/constants"
	"account-server/model"
	"account-server/service/dto"
	"account-server/utils"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

var userSerive *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userSerive == nil {
		userSerive = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userSerive
}

func SetLoginUserTokenToRedis(uId uint, token string) error {
	return global.RedisClient.Set(strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(uId)), -1), token, viper.GetDuration("jwt.tokenExpires")*time.Minute)
}

func GenerateTokenAndSetTokenToRedis(uId uint, userName string) (string, error) {
	token, err := utils.GenerateToken(uId, userName)
	if err == nil {
		err = SetLoginUserTokenToRedis(uId, token)
	}
	return token, err
}

func (u UserService) Login(IuserDTO dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token string
	iUser, err := u.Dao.GetUserByName(IuserDTO.Username)
	if err != nil || !utils.ComparePwd(iUser.Password, IuserDTO.Password) {
		errResult = errors.New("账号或密码不正确")
	} else {
		token, err = GenerateTokenAndSetTokenToRedis(iUser.ID, iUser.Username)
		if err != nil {
			errResult = errors.New(fmt.Sprintf("GenerateToken Error:%s", err.Error()))
		}
	}
	return iUser, token, errResult
}

func (u UserService) AddUser(iUserAddDTO dto.UserAddDTO) error {
	if u.Dao.CheckUserNameExist(iUserAddDTO.Username) {
		return errors.New("用户名已存在")
	}
	return u.Dao.AddUser(iUserAddDTO)
}

func (u UserService) GetUserById(iCommonIDDTO dto.CommonIDDTO) (model.User, error) {
	return u.Dao.GetUserById(iCommonIDDTO.ID)
}

func (u UserService) GetUserList(iUserListDTO dto.UserListDTO) ([]model.User, int64, error) {
	return u.Dao.GetUserList(iUserListDTO)
}

func (u UserService) EditUser(iUserEditDTO *dto.UserEditDTO) error {
	if iUserEditDTO.ID == 0 {
		return errors.New("invalid user ID")
	}
	return u.Dao.EditUser(iUserEditDTO)
}

func (u UserService) DelUser(iCommonIDDTO dto.CommonIDDTO) error {
	if iCommonIDDTO.ID == 0 {
		return errors.New("invalid user ID")
	}
	return u.Dao.DelUser(iCommonIDDTO.ID)
}
