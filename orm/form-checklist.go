package orm

import "github.com/FourWD/middleware/orm"

type FormChecklist struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name          string `json:"name" query:"name" gorm:"not null;type:varchar(500)"`
	CountOfDetail int    `json:"count_of_detail" query:"count_of_detail" gorm:"not null;type:int(2)"`
	orm.GormRowOrder
}
