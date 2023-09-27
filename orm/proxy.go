package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Proxy struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`

	UserID   string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)" firestore:"user_id"`
	Price    int       `json:"price" query:"price" gorm:"type:int" firestore:"price"`
	Created  time.Time `json:"created" query:"created" firestore:"created"`
	IsDelete bool      `json:"is_delete" query:"is_delete" gorm:"type:bool" firestore:"is_delete"`
	Deleted  time.Time `json:"deleted" query:"deleted" firestore:"deleted"`
}
