package orm

import "github.com/FourWD/auction.middleware/model"

type Vehicle struct {
	model.VehicleModel
	OpenPrice int `json:"open_price" query:"open_price" gorm:"type:int(10)"`
}
