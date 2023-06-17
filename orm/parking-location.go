package orm

import "github.com/FourWD/middleware/orm"

type ParkingLocation struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	ProvinceID string `json:"province_id" query:"province_id" db:"province_id" gorm:"type:varchar(36)"`
	Name       string `json:"name" query:"name" db:"name" gorm:"not null;type:varchar(50)"`
}
