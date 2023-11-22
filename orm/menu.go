package orm

import "github.com/FourWD/middleware/orm"

type Menu struct {
	orm.Menu
	UserTypeList string `json:"user_type_list" query:"user_type_list" gorm:"type:varchar(10)"`
}
