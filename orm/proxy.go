package orm

import (
	"time"
)

type Proxy struct {
	ID              string    `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	AuctionID       string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_price"`
	VehicleID       string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_price"`
	Price           int       `json:"price" query:"price" gorm:"type:int;uniqueIndex:idx_auction_id_vehicle_id_price"`
	UserID          string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	UserAuctionCode string    `json:"user_auction_code" query:"user_auction_code" gorm:"type:varchar(5)"`
	Created         time.Time `json:"created" query:"created"`
	CreatedUnixNano int64     `json:"created_unix_nano" query:"created_unix_nano" gorm:"type:int"`
	IsProcess       bool      `json:"is_process" query:"is_process" gorm:"type:bool"`
	ProcessDate     time.Time `json:"process_date" query:"process_date"`
	IsClear         bool      `json:"is_clear" query:"is_clear" gorm:"type:bool"`
	Remark          string    `json:"remark" query:"remark" gorm:"type:varchar(100)"`
	//IsWin           bool      `json:"is_win" query:"is_win" gorm:"type:bool"`
}
