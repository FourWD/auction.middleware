package orm

import "github.com/FourWD/middleware/model"

type AppInfo struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AppEnv     string `json:"app_env" query:"app_env" gorm:"type:varchar(20)"`
	AppVersion string `json:"app_version" query:"app_version" gorm:"type:varchar(20)"`
	AppUrl     string `json:"app_url" query:"app_url" gorm:"type:varchar(500)"`
}
