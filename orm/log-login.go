package orm

import "github.com/FourWD/middleware/orm"

type LogLogin struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID    string `json:"user_id" query:"user_id" gorm:"type:varchar(36);"`
	IsSuccess bool   `json:"is_success" query:"is_success" gorm:"bool"`
}
