package orm

import "github.com/FourWD/middleware/model"

type AppInfo struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AppVersion string `json:"app_version" query:"button_url" gorm:"type:varchar(20)"`
	AppUrl     string `json:"app_url" query:"app_url" gorm:"type:varchar(500)"`
	IsLasted   bool   `json:"is_lasted" query:"is_lasted" gorm:"type:bool"`
}
