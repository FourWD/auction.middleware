package orm

import "github.com/FourWD/middleware/model"

type AuctionVehicleBranch struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	BranchID     string `json:"branch_id" query:"branch_id" gorm:"type:varchar(36)"`
	TotalVehicle int    `json:"total_vehicle" query:"total_vehicle" gorm:"type:int"`
}
