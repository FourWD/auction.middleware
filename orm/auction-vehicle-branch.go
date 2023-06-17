package orm

import "github.com/FourWD/middleware/orm"

type AuctionVehicleBranch struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BranchID string `json:"branch_id" query:"branch_id" db:"branch_id" gorm:"type:varchar(36)"`

	TotalVehicle string `json:"total_vehicle" query:"total_vehicle" db:"total_vehicle" gorm:"type:int(11)"`
}
