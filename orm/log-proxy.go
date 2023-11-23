package orm

import (
	"time"
)

type LogProxy struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`

	AuctionID string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	FromPrice int    `json:"from_price" query:"from_price" gorm:"type:int"`
	ToPrice   int    `json:"to_price" query:"to_price" gorm:"type:int"`
	IsDelete  bool   `json:"is_delete" query:"is_delete" gorm:"type:bool"`

	UserID string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`

	Created          time.Time `json:"created" query:"created"`
	CreatedUnixMilli int64     `json:"created_unix_milli" query:"created_unix_milli" gorm:"type:int"`
}
