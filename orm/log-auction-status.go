package orm

import "github.com/FourWD/middleware/model"

type LogAuctionStatus struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AuctionID       string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
	AuctionStatusID string `json:"auction_status_id" query:"auction_status_id" gorm:"type:varchar(2)"`
}
