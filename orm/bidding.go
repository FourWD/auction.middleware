package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Bidding struct {
	ID uint `json:"id" query:"id" gorm:"primaryKey;autoIncrement"`
	orm.GormModel

	AuctionID string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_bidding"`
	VehicleID string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_bidding"`
	UserID    string    `json:"user_id" query:"user_id" gorm:"type:varchar(36);uniqueIndex:idx_bidding" firestore:"user_id" `
	Price     int       `json:"price" query:"price" gorm:"type:int" firestore:"price" `
	IsProxy   bool      `json:"is_proxy" query:"is_proxy" gorm:"type:bool" firestore:"is_proxy"`
	Created   time.Time `json:"created" query:"created" firestore:"created"`
}
