package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type Employee struct {
	midOrm.Employee
	SourceList  string `json:"source_list" query:"source_list" gorm:"type:varchar(100)"`
	FinanceList string `json:"finance_list" query:"finance_list" gorm:"type:varchar(100)"`
}
