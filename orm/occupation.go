package orm

import "github.com/FourWD/middleware/orm"

type Occupation struct {
	orm.Occupation
	FinanceID    string `json:"finance_id" query:"finance_id" gorm:"type:varchar(10)"`
	BusinessType string `json:"business_type" query:"business_type" gorm:"type:varchar(2)"`
}
