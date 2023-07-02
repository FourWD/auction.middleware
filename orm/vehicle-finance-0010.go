package orm

import "github.com/FourWD/middleware-auction/model"

type VehicleFinance0010 struct {
	model.VehicleModel
}

func (VehicleFinance0010) TableName() string {
	return "vehicle-finances-0010"
}
