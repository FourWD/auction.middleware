package orm

import "github.com/FourWD/middleware/orm"

type TermsCondition struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	TCVersion   string `json:"tc_version" query:"tc_version" gorm:"type:varchar(100)"`
	Description string `json:"description" query:"description" gorm:"type:text"`
}
