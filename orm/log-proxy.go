package orm

import (
	"time"
)

type LogProxy struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`

	AuctionID       string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
	VehicleID       string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	UserID          string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	Price           int       `json:"price" query:"price" gorm:"type:int"`
	NewPrice        int       `json:"new_price" query:"new_price" gorm:"type:int"`
	LogProxyTypeID  string    `json:"log_proxy_type_id" query:"log_proxy_type_id" gorm:"type:varchar(1)"`
	Created         time.Time `json:"created" query:"created"`
	CreatedUnixNano int64     `json:"created_unix_nano" query:"created_unix_nano" gorm:"type:int"`
}
