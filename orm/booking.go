package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type Booking struct {
	midOrm.Booking
	AuctionID string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
}
