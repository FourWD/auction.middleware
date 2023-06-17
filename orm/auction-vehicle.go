package orm

import "github.com/FourWD/middleware/orm"

type AuctionVehicle struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID string `json:"finance_id" query:"finance_id" db:"finance_id" gorm:"type:varchar(36)"`

	UserID string `json:"user_id" query:"user_id" db:"user_id" gorm:"type:varchar(36)"`
}
