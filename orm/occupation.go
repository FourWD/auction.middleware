package orm

import "github.com/FourWD/middleware/orm"

type Occupation struct {
	orm.Occupation
	FinanceID string `json:"finance_id" query:"finance_id" gorm:"type:varchar(10)"`
}
