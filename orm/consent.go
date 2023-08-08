package orm

import "github.com/FourWD/middleware/orm"

type Consent struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	ConsentVersion string `json:"consent_version" query:"consent_version" db:"consent_version" gorm:"type:varchar(3)"`
	Detail         string `json:"detail" query:"detail" db:"detail" gorm:"type:varchar(500)"`
}
