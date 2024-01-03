package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type UserAcceptMinPrice struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	UserID    string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	AuctionID string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`

	AcceptDate time.Time `json:"accept_date" query:"accept_date"`
}

// ทำเส้น check min price กับ save ลง tb นี้ ขึ้น
