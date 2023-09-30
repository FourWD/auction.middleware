package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type AuctionVehicleUserFavorite struct {
	ID uint `json:"id" query:"id" gorm:"primaryKey;autoIncrement"`
	orm.GormModel

	AuctionID       string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_user_id"`
	UserID          string    `json:"user_id" query:"user_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_user_id"`
	VehicleID       string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id_user_id"`
	UserAuctionCode string    `json:"user_auction_code" query:"user_auction_code" gorm:"type:varchar(15)"`
	IsJoin          bool      `json:"is_join" query:"is_join" gorm:"bool"`
	JoinDate        time.Time `json:"join_date" query:"join_date"`
	IsFavorite      bool      `json:"is_favorite" query:"is_favorite" gorm:"bool"`
	FavoriteDate    time.Time `json:"favorite_date" query:"favorite_date"`
	IsBidding       bool      `json:"is_bidding" query:"is_bidding" gorm:"bool"`
	MD5             string    `json:"md5" query:"md5" gorm:"type:varchar(32)"`
}
