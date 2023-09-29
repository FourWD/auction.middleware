package orm

import "github.com/FourWD/auction.middleware/model"

type VehicleSource0001 struct {
	model.VehicleModel
}

func (VehicleSource0001) TableName() string {
	return "vehicle-sources-0001"
}
