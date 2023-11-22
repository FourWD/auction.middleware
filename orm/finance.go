package orm

import "github.com/FourWD/middleware/orm"

type Finance struct {
	orm.Finance
	IsRegisterActive bool `json:"is_register_active" query:"is_register_active" gorm:"type:bool"`
	ValidDay         int  `json:"valid_day" query:"valid_day" gorm:"type:int"`
}
