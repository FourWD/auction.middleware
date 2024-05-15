package orm

import (
	"time"
)

type AuctionUser struct {
	// ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	// model.GormModel

	UserID       string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	RightDeposit int       `json:"right_deposit" query:"right_deposit" gorm:"type:int(5)"`
	EodDate      time.Time `json:"eod_date" query:"eod_date"`
}
