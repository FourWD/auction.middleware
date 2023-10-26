package orm

import "github.com/FourWD/middleware/orm"

type VehicleDetail struct {
	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SoldStatusID string `json:"sold_status_id" query:"sold_status_id" gorm:"type:varchar(2)"`
}
