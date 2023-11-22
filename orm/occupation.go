package orm

import (
	"github.com/FourWD/middleware/model"
)

type Ocupation struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel
	Name        string `json:"name" query:"name" gorm:"type:varchar(300)"`
	FinanceID   string `json:"finance_id" query:"finance_id" gorm:"type:varchar(10)"`
	OcupationID string `json:"ocupation_id" query:"ocupation_id" gorm:"type:varchar(36)"`
	model.GormRowOrder
}
