package orm

import "github.com/FourWD/middleware/orm"

type AuctionParkingLocation struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	ParkingLocationID string `json:"parking_location_id" query:"parking_location_id" db:"parking_location_id" gorm:"type:varchar(36)"`

	TotalVehicle string `json:"total_vehicle" query:"total_vehicle" db:"total_vehicle" gorm:"type:int(11)"`
}
