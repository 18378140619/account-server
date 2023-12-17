package dao

import (
	"account-server/service/dto"
	"gorm.io/gorm"
)

// 通用分页

func Paginate(p dto.Paginate) func(orm *gorm.DB) *gorm.DB {
	return func(orm *gorm.DB) *gorm.DB {
		return orm.Offset((p.GetPage() - 1) * p.GetLimit()).Limit(p.GetLimit())
	}
}
