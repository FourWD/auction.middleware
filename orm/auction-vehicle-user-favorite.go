package orm

import "github.com/FourWD/middleware/orm"

type AuctionVehicleUserFavorite struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID string `json:"finance_id" query:"finance_id" gorm:"type:varchar(36)"`

	UserID       string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	VehicleID    string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	IsJoin       bool   `json:"is_join" query:"is_join" gorm:"bool"`
	JoinDate     string `json:"join_date" query:"join_date" gorm:"not null;type:varchar(15)"`
	IsFavorite   bool   `json:"is_favorite" query:"is_favorite" gorm:"bool"`
	FavoriteDate string `json:"favorite_date" query:"favorite_date" gorm:"not null;type:varchar(15)"`
	IsBidding    bool   `json:"is_bidding" query:"is_bidding" gorm:"bool"`
}
