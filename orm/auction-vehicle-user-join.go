package orm

import "github.com/FourWD/middleware/orm"

type AuctionVehicleUserFavorite struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID string `json:"finance_id" query:"finance_id" db:"finance_id" gorm:"type:varchar(36)"`

	UserID           string `json:"user_id" query:"user_id" db:"user_id" gorm:"type:varchar(36)"`
	VehicleTableCode string `json:"vehicle_table_code" query:"vehicle_table_code" db:"vehicle_table_code" gorm:"type:varchar(4)"`
	VehicleID        string `json:"vehicle_id" query:"vehicle_id" db:"vehicle_id" gorm:"type:varchar(36)"`
	IsJoin           bool   `json:"is_join" query:"is_join" db:"is_join" gorm:"boolean"`
	JoinDate         string `json:"join_date" query:"join_date" db:"join_date" gorm:"not null;type:varchar(15)"`
	IsFavorite       bool   `json:"is_favorite" query:"is_favorite" db:"is_favorite" gorm:"boolean"`
	FavoriteDate     string `json:"favorite_date" query:"favorite_date" db:"favorite_date" gorm:"not null;type:varchar(15)"`
	IsBidding        bool   `json:"is_bidding" query:"is_bidding" db:"is_bidding" gorm:"boolean"`
}
