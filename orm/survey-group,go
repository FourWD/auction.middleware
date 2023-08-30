package orm

import "github.com/FourWD/middleware/orm"

type SurveyGroup struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Subject string `json:"name" query:"name" db:"name" gorm:"not null;type:varchar(100)"`
	Remark  string `json:"remark" query:"remark" db:"name" gorm:"type:text;"`
}
