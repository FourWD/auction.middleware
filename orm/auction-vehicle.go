package orm

import "github.com/FourWD/middleware/orm"

type AuctionVehicle struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID string `json:"finance_id" query:"finance_id" db:"finance_id" gorm:"type:varchar(36)"`
	VehicleID string `json:"vehicle_id" query:"vehicle_id" db:"vehicle_id" gorm:"type:varchar(36)"`
}
