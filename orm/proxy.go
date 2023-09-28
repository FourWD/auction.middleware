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

	UserID           string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	Price            int       `json:"price" query:"price" gorm:"type:int"`
	Created          time.Time `json:"created" query:"created"`
	CreatedUnixMilli int64     `json:"created_unix_milli" query:"created_unix_milli" gorm:"type:int"`
	IsDelete         bool      `json:"is_delete" query:"is_delete" gorm:"type:bool"`
	Deleted          time.Time `json:"deleted" query:"deleted"`
	DeletedUnixMilli int64     `json:"deleted_unix_milli" query:"deleted_unix_milli" gorm:"type:int"`
}
