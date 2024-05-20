package orm

import (
	"github.com/FourWD/middleware/model"
)

type LogAuctionVehicle struct {
	ID string `json:"id" query:"id" firestore:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AuctionID       string `json:"auction_id" query:"auction_id" firestore:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`
	VehicleID       string `json:"vehicle_id" query:"vehicle_id" firestore:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`
	AuctionStatusID string `json:"auction_status_id" query:"auction_status_id" firestore:"auction_status_id" gorm:"type:varchar(2)"`
}
