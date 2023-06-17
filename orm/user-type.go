package orm

import "github.com/FourWD/middleware/orm"

type UserType struct { // 01 Seller 02 Customer
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name string `json:"name" query:"name" db:"name" gorm:"type:varchar(20);"`
	orm.RowOrder
}
