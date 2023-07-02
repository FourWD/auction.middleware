package orm

import "github.com/FourWD/middleware-auction/model"

type VehicleFinance0007 struct {
	model.VehicleModel
}

func (VehicleFinance0007) TableName() string {
	return "vehicle-finances-0007"
}
