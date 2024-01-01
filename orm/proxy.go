package orm

import (
	"time"
)

type Proxy struct {
	ID               string    `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	AuctionID        string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_price"`
	VehicleID        string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_price"`
	Price            int       `json:"price" query:"price" gorm:"type:int;uniqueIndex:idx_auction_id_vehicle_id_price"`
	UserID           string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	UserAuctionCode  string    `json:"user_auction_code" query:"user_auction_code" gorm:"type:varchar(5)"`
	Created          time.Time `json:"created" query:"created"`
	CreatedUnixMilli int64     `json:"created_unix_milli" query:"created_unix_milli" gorm:"type:int"`
	IsProcess        bool      `json:"is_process" query:"is_process" gorm:"type:bool"`
	ProcessDate      time.Time `json:"process_data" query:"process_data"`
}
