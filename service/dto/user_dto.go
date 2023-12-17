package dto

import (
	"account-server/model"
)

type UserLoginDTO struct {
	Username string `json:"Username" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

//===============
//用户先关DTO

type UserAddDTO struct {
	ID       uint
	Username string `json:"username" form:"username" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" form:"real_name"`
	Email    string `json:"email" form:"email"`
	Mobile   string `json:"mobile" form:"mobile"`
	Avatar   string
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"密码不能为空"`
}

type UserEditDTO struct {
	ID       uint   `json:"id" binding:"required" message:"ID不能为空"`
	Username string `json:"username" form:"username"`
	RealName string `json:"real_name" form:"real_name"`
	Email    string `json:"email" form:"email"`
	Mobile   string `json:"mobile" form:"mobile"`
	Avatar   string
}

func (u *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Username = u.Username
	iUser.RealName = u.RealName
	iUser.Email = u.Email
	iUser.Mobile = u.Mobile
	iUser.Password = u.Password
}

func (u *UserEditDTO) ConvertToModel(iUser *model.User) {
	iUser.Username = u.Username
	iUser.ID = u.ID
	iUser.Email = u.Email
	iUser.Mobile = u.Mobile
	iUser.RealName = u.RealName
}

// UserListDTO ==============
// 用户列表DTO
type UserListDTO struct {
	Paginate
}
