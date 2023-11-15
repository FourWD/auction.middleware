package orm

import "github.com/FourWD/middleware/orm"

type VehicleSubType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleTypeID string `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(10);"`

	Name string `json:"name" query:"name" gorm:"type:varchar(100);"`
	orm.GormRowOrder
}
