package orm

import "github.com/FourWD/middleware/orm"

type AuctionVehicle struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID        string `json:"finance_id" query:"finance_id" db:"finance_id" gorm:"type:varchar(36)"`
	VehicleTableCode string `json:"vehicle_table_code" query:"vehicle_table_code" db:"vehicle_table_code" gorm:"type:varchar(4)"`
	VehicleID        string `json:"vehicle_id" query:"vehicle_id" db:"vehicle_id" gorm:"type:varchar(36)"`

	VehicleNo string `json:"vehicle_no" query:"vehicle_no" db:"vehicle_no" gorm:"type:varchar(10);"`

	BiddingStep1 int `json:"bidding_step_1" query:"bidding_step_1" db:"bidding_step_1" gorm:"type:int(6);"`
	BiddingStep2 int `json:"bidding_step_2" query:"bidding_step_2" db:"bidding_step_2" gorm:"type:int(6);"`
	BiddingStep3 int `json:"bidding_step_3" query:"bidding_step_3" db:"bidding_step_3" gorm:"type:int(6);"`
}
