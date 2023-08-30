package orm

import "github.com/FourWD/middleware/orm"

type VehicleFueltype struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleTypeID  string `json:"vehicle_type_id" query:"vehicle_type_id" db:"vehicle_type_id" gorm:"type:varchar(36);"`
	VehicleBrandID string `json:"vehicle_brand_id" query:"vehicle_brand_id" db:"vehicle_brand_id" gorm:"type:varchar(36);"`
	VehicleGearID  string `json:"vehicle_gear_id" query:"vehicle_gear_id" db:"vehicle_gear_id" gorm:"type:varchar(36);"`
	Name           string `json:"name" query:"name" db:"name" gorm:"type:varchar(20);"`
	Seat           string `json:"seat" query:"seat" db:"seat" gorm:"type:int(1);"`
	EngineCC       int    `json:"engine_cc" query:"engine_cc" db:"engine_cc" gorm:"type:varchar(4);"`
	orm.RowOrder
}
