package model

import (
	"github.com/FourWD/middleware/orm"
)

type VehicleModel struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	License   string `json:"license" query:"license" db:"license" gorm:"type:varchar(20);"`
	OpenPrice string `json:"open_price" query:"open_price" db:"open_price" gorm:"type:varchar(20);"`
	MinPrice  string `json:"min_price" query:"min_price" db:"min_price" gorm:"type:varchar(20);"`
	ParkingLocationID string `json:"parking_location_id" query:"parking_location_id" db:"parking_location_id" gorm:"type:varchar(36)"`
}
