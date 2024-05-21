package orm

import (
	"github.com/FourWD/middleware/model"
)

type LogFavorite struct {
	ID string `json:"id" query:"id" firestore:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AuctionID string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	UserID    string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
}
