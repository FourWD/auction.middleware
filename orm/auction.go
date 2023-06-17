package orm

import "github.com/FourWD/middleware/orm"

type Auction struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FinanceID string `json:"finance_id" query:"finance_id" db:"finance_id" gorm:"type:varchar(36)"`

	StartDate string `json:"start_date" query:"start_date" db:"start_date" gorm:"not null;type:varchar(15)"`
	EndDate   string `json:"end_date" query:"end_date" db:"end_date" gorm:"not null;type:varchar(15)"`
	IsShow    bool   `json:"is_show" query:"is_show" db:"is_show" gorm:"boolean"`
}
