package orm

import "github.com/FourWD/middleware/orm"

type VehicleOpenPrice struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name      string `json:"name" query:"name" gorm:"type:varchar(20);"`
	OpenPrice string `json:"open_price" query:"open_price" gorm:"type:varchar(10);"`
	EndPrice  string `json:"end_price" query:"end_price" gorm:"type:varchar(10);"`
	orm.RowOrder
}
