package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type AuctionVehicle struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`
	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`

	VehicleNo string `json:"vehicle_no" query:"vehicle_no" gorm:"type:varchar(10); "`

	OpenPrice  int       `json:"open_price" query:"open_price" gorm:"type:int(10);"`
	MinPrice   int       `json:"min_price" query:"min_price" gorm:"type:int(10);"`
	ClosePrice int       `json:"close_price" query:"close_price" gorm:"type:int(10);"`
	EndDate    time.Time `json:"end_date" query:"end_date"`
	IsExtra    bool      `json:"is_extra" query:"is_extra" gorm:"bool"`
	IsEnd      bool      `json:"is_end" query:"is_end" gorm:"bool"`

	BiddingStep1 int `json:"bidding_step_1" query:"bidding_step_1" gorm:"column:bidding_step_1;type:int(6);"`
	BiddingStep2 int `json:"bidding_step_2" query:"bidding_step_2" gorm:"column:bidding_step_2;type:int(6);"`
	BiddingStep3 int `json:"bidding_step_3" query:"bidding_step_3" gorm:"column:bidding_step_3;type:int(6);"`
}
