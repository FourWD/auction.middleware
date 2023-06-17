package orm

import "github.com/FourWD/middleware/orm"

type Branch struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Label      string  `json:"label" query:"label" db:"label" gorm:"type:varchar(100)"`
	Name       string  `json:"name" query:"name" db:"name" gorm:"type:varchar(100)"`
	Address    string  `json:"address" query:"address" db:"address" gorm:"type:text"`
	ProvinceID string  `json:"province_id" query:"province_id" db:"province_id" gorm:"type:varchar(36)"`
	Phone1     string  `json:"phone_1" query:"phone_1" db:"phone_1" gorm:"type:varchar(20)"`
	Phone2     string  `json:"phone_2" query:"phone_2" db:"phone_2" gorm:"type:varchar(20)"`
	Line       string  `json:"line" query:"line" db:"line" gorm:"type:varchar(20)"`
	Facebook   string  `json:"facebook" query:"facebook" db:"facebook" gorm:"type:varchar(20)"`
	Tiktok     string  `json:"tiktok" query:"tiktok" db:"tiktok" gorm:"type:varchar(20)"`
	Latitude   float64 `json:"latitude" query:"latitude" db:"latitude" gorm:"type:float"`
	Longitude  float64 `json:"longitude" query:"longitude" db:"longitude" gorm:"type:float"`
	orm.RowOrder
}
