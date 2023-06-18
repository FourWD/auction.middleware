package orm

import "middleware-auction/model"

type VehicleFinance0006 struct {
	model.VehicleModel
}

func (VehicleFinance0006) TableName() string {
	return "vehicle-finances-0006"
}
