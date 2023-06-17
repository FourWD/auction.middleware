package orm

import "github.com/FourWD/middleware/orm"

type TeamCondition struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	TCVersion   string `json:"tc_version" query:"tc_version" db:"tc_version" gorm:"type:varchar(100)"`
	Description string `json:"description" query:"description" db:"description" gorm:"type:text"`
}
