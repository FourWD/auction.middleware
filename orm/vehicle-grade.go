package orm

import "github.com/FourWD/middleware/orm"

type VehicleGrade struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name        string `json:"name" query:"name" db:"name" gorm:"type:varchar(20);"`
	Description string `json:"description" query:"description" db:"description" gorm:"type:varchar(20);"`
	orm.RowOrder
}
