package orm

import "github.com/FourWD/middleware/orm"

type UserType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name string `json:"name" query:"name" gorm:"type:varchar(20);"`
	orm.RowOrder
}

/*
บุคคลธรรมดา นิติบุคคล
*/
