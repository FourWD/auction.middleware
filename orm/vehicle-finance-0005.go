package orm

import "middleware-auction/model"

type VehicleFinance0005 struct {
	model.VehicleModel
}

func (VehicleFinance0005) TableName() string {
	return "vehicle-finances-0005"
}
