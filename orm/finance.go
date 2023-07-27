package orm

import "github.com/FourWD/middleware/orm"

type Finance struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Logo    string `json:"logo" query:"logo" db:"logo" gorm:"type:varchar(50)"`
	Code    string `json:"code" query:"code" db:"code" gorm:"type:varchar(4)"`
	Name    string `json:"name" query:"name" db:"name" gorm:"type:varchar(50)"`
	Color   string `json:"color" query:"color" db:"color" gorm:"type:varchar(7)"`
	LoanUrl string `json:"loan_url" query:"loan_url" db:"loan_url" gorm:"type:varchar(200)"`
	orm.RowOrder
}
