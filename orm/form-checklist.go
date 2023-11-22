package orm

import (
	"github.com/FourWD/middleware/model"
)

type FormChecklist struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	Name          string `json:"name" query:"name" gorm:"not null;type:varchar(500)"`
	CountOfDetail int    `json:"count_of_detail" query:"count_of_detail" gorm:"not null;type:int(2)"`
	model.GormRowOrder
}
