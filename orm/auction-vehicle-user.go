package orm

import (
	"time"
)

type AuctionVehicleUser struct {
	ID uint `json:"id" query:"id" gorm:"primaryKey;autoIncrement"`

	AuctionID       string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_user_id"`
	UserID          string    `json:"user_id" query:"user_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_user_id"`
	VehicleID       string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_user_id"`
	UserAuctionCode string    `json:"user_auction_code" query:"user_auction_code" gorm:"type:varchar(5)"`
	IsJoin          bool      `json:"is_join" query:"is_join" gorm:"bool"`
	JoinDate        time.Time `json:"join_date" query:"join_date"`
	IsFavorite      bool      `json:"is_favorite" query:"is_favorite" gorm:"bool"`
	FavoriteDate    time.Time `json:"favorite_date" query:"favorite_date"`
	IsBidding       bool      `json:"is_bidding" query:"is_bidding" gorm:"bool"`
	ProxyID         string    `json:"proxy_id" query:"proxy_id" gorm:"type:varchar(36)"`
	CountView       int       `json:"count_view" query:"count_view" gorm:"type:int"`
}
