package orm

import "github.com/FourWD/middleware-auction/model"

type VehicleSource0002 struct {
	model.VehicleModel
}

func (VehicleSource0002) TableName() string {
	return "vehicle-sources-0002"
}