package orm

import "github.com/FourWD/middleware-auction/model"

type VehicleSource0003 struct {
	model.VehicleModel
}

func (VehicleSource0003) TableName() string {
	return "vehicle-sources-0003"
}
