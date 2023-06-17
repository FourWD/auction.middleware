package orm

import "github.com/FourWD/middleware/orm"

type Province struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name string `json:"name" query:"name" db:"name" gorm:"not null;type:varchar(100)"`
}
