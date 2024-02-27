package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type Employee struct {
	midOrm.Employee
	SourceList string `json:"source_list" query:"source_list" gorm:"type:varchar(100)"`
}
