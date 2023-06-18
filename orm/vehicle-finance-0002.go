package orm

import "middleware-auction/model"

type VehicleFinance0002 struct {
	model.VehicleModel
	Name string `gorm:"table:vehicle-finances-0002"`
}
