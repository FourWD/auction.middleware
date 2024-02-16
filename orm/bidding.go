package orm

import (
	"time"
)

// type Bidding struct {
// 	ID uint `json:"id" query:"id" gorm:"primaryKey;autoIncrement"`
// 	model.GormModel

// 	AuctionID       string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_bidding"`
// 	VehicleID       string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_bidding"`
// 	UserID          string    `json:"user_id" query:"user_id" gorm:"type:varchar(36);uniqueIndex:idx_bidding"`
// 	Price           int       `json:"price" query:"price" gorm:"type:int"`
// 	BiddingStep     int       `json:"bidding_step" query:"bidding_step" gorm:"type:int"`
// 	IsProxy         bool      `json:"is_proxy" query:"is_proxy" gorm:"type:bool"`
// 	Created         time.Time `json:"created" query:"created"`
// 	CreatedUnixNano int64     `json:"created_unix_nano" query:"created_unix_nano" gorm:"type:int"`
// }

type Bidding struct {
	ID              string    `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	AuctionID       string    `json:"auction_id" query:"auction_id" firestore:"auction_id"`
	VehicleID       string    `json:"vehicle_id" query:"vehicle_id" firestore:"vehicle_id"`
	UserID          string    `json:"user_id" query:"user_id" firestore:"user_id"`
	UserAuctionCode string    `json:"user_auction_code" query:"user_auction_code" firestore:"user_auction_code"`
	UserDisplayName string    `json:"user_display_name" query:"user_id" firestore:"user_display_name"`
	Price           int       `json:"price" query:"price" firestore:"price"`
	BiddingStep     int       `json:"bidding_step" query:"bidding_step" firestore:"bidding_step"`
	IsProxy         bool      `json:"is_proxy" query:"is_proxy" firestore:"is_proxy"`
	ProxyID         string    `json:"proxy_id" query:"proxy_id" firestore:"proxy_id"`
	IsOpenPrice     bool      `json:"is_open_price" query:"is_open_price" firestore:"is_open_price"`
	RowOrder        int       `json:"row_order" query:"row_order" firestore:"row_order"`
	Created         time.Time `json:"created" query:"created" firestore:"created"`
	CreatedUnixNano string    `json:"created_unix_nano" query:"created_unix_nano" firestore:"created_unix_nano"` //  larger than firebase maximum integer
}
