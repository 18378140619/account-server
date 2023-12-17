package dao

var loginDao *LoginDao

type LoginDao struct {
	BaseDao
}

func NewLoginDao() *LoginDao {
	if loginDao == nil {
		loginDao = &LoginDao{
			BaseDao: NewBaseDao(),
		}
	}
	return loginDao
}
