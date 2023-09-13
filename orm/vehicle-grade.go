package orm

import "github.com/FourWD/middleware/orm"

type VehicleGrade struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name        string `json:"name" query:"name" db:"name" gorm:"type:varchar(50);"`
	Description string `json:"description" query:"description" gorm:"type:varchar(50);"`
	orm.RowOrder
}
