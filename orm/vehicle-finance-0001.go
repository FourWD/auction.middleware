package orm

import "middleware-auction/model"

type VehicleFinance0001 struct {
	model.VehicleModel
	Name string `gorm:"table:vehicle-finances-0001"`
}
