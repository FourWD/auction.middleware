package orm

import "github.com/FourWD/middleware/model"

type VehicleDetail struct {
	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	VehicleSoldStatusID string `json:"vehicle_sold_status_id" query:"vehicle_sold_status_id" gorm:"type:varchar(2)"`
}
