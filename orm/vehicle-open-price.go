package orm

import (
	"github.com/FourWD/middleware/model"
)

type VehicleOpenPrice struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key"`
	model.GormModel

	Name      string `json:"name" query:"name" gorm:"type:varchar(20)"`
	OpenPrice int    `json:"open_price" query:"open_price" gorm:"type:int"`
	EndPrice  int    `json:"end_price" query:"end_price" gorm:"type:int"`
	model.GormRowOrder
}
