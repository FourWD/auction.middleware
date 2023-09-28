package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Bidding struct {
	ID uint `json:"id" query:"id" gorm:"primaryKey;autoIncrement"`
	orm.GormModel

	AuctionID        string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_bidding"`
	VehicleID        string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_bidding"`
	UserID           string    `json:"user_id" query:"user_id" gorm:"type:varchar(36);uniqueIndex:idx_bidding"`
	Price            int       `json:"price" query:"price" gorm:"type:int"`
	BiddingStep      int       `json:"bidding_step" query:"bidding_step" gorm:"type:int"`
	IsProxy          bool      `json:"is_proxy" query:"is_proxy" gorm:"type:bool"`
	Created          time.Time `json:"created" query:"created"`
	CreatedUnixMilli int64     `json:"created_unix_milli" query:"created_unix_milli" gorm:"type:int"`
}
