package orm

import "github.com/FourWD/middleware/orm"

type VehicleType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name   string `json:"name" query:"name" gorm:"type:varchar(20);"`
	NameTh string `json:"name_th" query:"name_th" gorm:"type:varchar(20);"`
	orm.GormRowOrder
}
