package orm

import "github.com/FourWD/middleware/orm"

type VehicleSubModel struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleModelID string `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36);"`

	VehicleDrivetypeID string `json:"vehicle_drivetype_id" query:"vehicle_drivetype_id" gorm:"type:varchar(2);"`
	VehicleGearID      string `json:"vehicle_gear_id" query:"vehicle_gear_id" gorm:"type:varchar(2);"`
	VehicleFuelID      string `json:"vehicle_fuel_id" query:"vehicle_fuel_id" gorm:"type:varchar(2);"`
	Name               string `json:"name" query:"name" gorm:"type:varchar(50);"`
	Seat               string `json:"seat" query:"seat" gorm:"type:int(1);"`
	EngineCC           int    `json:"engine_cc" query:"engine_cc" gorm:"type:varchar(4);"`
	orm.RowOrder
}
