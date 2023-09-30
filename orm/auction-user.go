package orm

import "github.com/FourWD/middleware/orm"

type AuctionUser struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID string `json:"finance_id" query:"finance_id" gorm:"type:varchar(36)"`
	UserID       string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	RightFinance int    `json:"right_finance" query:"right_finance" gorm:"type:int(5)"`
}
