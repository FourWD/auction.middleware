package orm

import "github.com/FourWD/middleware/orm"

type UserDevice struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID         string `json:"user_id" query:"user_id" gorm:"type:varchar(36);"`
	DeviceID       string `json:"device_id" query:"device_id" gorm:"type:varchar(36);"`
	DeviceName     string `json:"device_name" query:"device_name" gorm:"type:varchar(50);"`
	DeviceLocation string `json:"device_lcation" query:"device_lcation" gorm:"type:varchar(50);"`
	IsActive       int    `json:"is_active" query:"is_active" gorm:"type:int;"`
	LastActive     string `json:"last_active" query:"last_active" gorm:"type:varchar(50);"`
}

/*
บุคคลธรรมดา นิติบุคคล
*/
