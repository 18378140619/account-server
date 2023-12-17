package dao

import (
	"account-server/model"
	"account-server/service/dto"
	"fmt"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			BaseDao: NewBaseDao(),
		}
	}
	return userDao
}
func (u *UserDao) GetUserByName(stUserName string) (model.User, error) {
	var iUser model.User
	err := u.Orm.Model(&iUser).Where("username=?", stUserName).Find(&iUser).Error
	return iUser, err
}
func (u *UserDao) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var iUser model.User
	u.Orm.Model(&iUser).Where("username=? and password=?", stUserName, stPassword).Find(&iUser)
	return iUser
}

func (u *UserDao) AddUser(iUserAddDTO dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)
	err := u.Orm.Save(&iUser).Error
	if err == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}
	return err
}

func (u *UserDao) CheckUserNameExist(stUserName string) bool {
	var nTotal int64
	u.Orm.Model(&model.User{}).Where("username=?", stUserName).Count(&nTotal)
	return nTotal > 0
}

func (u *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User
	err := u.Orm.First(&iUser, id).Error
	return iUser, err
}

func (u *UserDao) GetUserList(iUserListDTO dto.UserListDTO) ([]model.User, int64, error) {
	var iUserList []model.User
	var nTotal int64
	err := u.Orm.Model(&model.User{}).Scopes(Paginate(iUserListDTO.Paginate)).
		Find(&iUserList).Offset(-1).Limit(-1).Count(&nTotal).Error
	return iUserList, nTotal, err
}

func (u *UserDao) EditUser(iUserEditDTO *dto.UserEditDTO) error {
	var iUser model.User
	u.Orm.First(&iUser, iUserEditDTO.ID)
	iUserEditDTO.ConvertToModel(&iUser)
	fmt.Printf("%+v", &iUser)
	return u.Orm.Save(&iUser).Error
}

func (u *UserDao) DelUser(id uint) error {
	return u.Orm.Delete(&model.User{}, id).Error
}
