package orm

import "github.com/FourWD/middleware/orm"

type LogClearWrongLogin struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	EmployeeID string `json:"employee_id" query:"employee_id" gorm:"type:varchar(36);"`
	UserID     string `json:"user_id" query:"user_id" gorm:"type:varchar(36);"`
}
