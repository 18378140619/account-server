package model

import (
	"account-server/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:64;not null"`
	Password string `json:"password" gorm:"size:128;not null"`
	RealName string `json:"realName" gorm:"size:128"`
	Email    string `json:"email" gorm:"size:128"`
	Mobile   string `json:"mobile" gorm:"size:128"`
	Avatar   string `json:"avatar" gorm:"size:11"`
}

func (u *User) Encrypt() error {
	hash, err := utils.Encrypt(u.Password)
	if err == nil {
		u.Password = hash
	}
	return err
}

func (u *User) BeforeCreate(orm *gorm.DB) (err error) {
	return u.Encrypt()
}

// =======
// = 用户登录信息
type LoginUser struct {
	ID       uint
	Username string
}
