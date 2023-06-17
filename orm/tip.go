package orm

import "github.com/FourWD/middleware/orm"

type Tip struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	RowOrder int `json:"row_order" query:"row_order" db:"row_order" gorm:"type:int(4);"`
}
