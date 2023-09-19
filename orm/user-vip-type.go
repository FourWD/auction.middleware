package orm

import "github.com/FourWD/middleware/orm"

type UserVipType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name string `json:"name" query:"name" gorm:"not null;type:varchar(50)"`
}

/*
0 success 1 reject
*/
