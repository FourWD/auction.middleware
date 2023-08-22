package orm

import "github.com/FourWD/middleware/orm"

type Tip struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
}
