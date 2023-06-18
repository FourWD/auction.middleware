package orm

import "middleware-auction/model"

type VehicleFinance0001 struct {
	model.VehicleModel
}

func (VehicleFinance0001) TableName() string {
	return "table:vehicle-finances-0001"
}
