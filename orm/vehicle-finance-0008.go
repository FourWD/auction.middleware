package orm

import "middleware-auction/model"

type VehicleFinance0008 struct {
	model.VehicleModel
}

func (VehicleFinance0008) TableName() string {
	return "vehicle-finances-0008"
}
