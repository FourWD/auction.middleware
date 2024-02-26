package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type Employee struct {
	midOrm.Employee
	SourceID string `json:"source_id" query:"source_id" gorm:"type:varchar(100)"`
}
