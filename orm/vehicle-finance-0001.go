package orm

import "github.com/FourWD/middleware-auction/model"

type VehicleFinance0001 struct {
	model.VehicleModel
}

func (VehicleFinance0001) TableName() string {
	return "vehicle-finances-0001"
}
