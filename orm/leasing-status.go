package orm

import "github.com/FourWD/middleware/orm"

type LeasingStatus struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name string `json:"name" query:"name" db:"name" gorm:"type:varchar(100)"`
	orm.RowOrder
}
