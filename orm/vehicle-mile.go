package orm

import "github.com/FourWD/middleware/orm"

type VehicleMile struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name    string `json:"name" query:"name" gorm:"type:varchar(20);"`
	MileMin string `json:"mile_min" query:"mile_min" gorm:"type:varchar(10);"`
	MileMax string `json:"mile_max" query:"mile_max" gorm:"type:varchar(10);"`
	orm.RowOrder
}
