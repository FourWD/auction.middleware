package orm

import "github.com/FourWD/auction.middleware/model"

type Vehicle struct {
	model.VehicleModel
}

func (Vehicle) TableName() string {
	return "vehicles"
}
