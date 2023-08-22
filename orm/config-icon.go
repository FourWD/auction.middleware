package orm

import "github.com/FourWD/middleware/orm"

type ConfigIcon struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name     string `json:"name" query:"name" gorm:"type:varchar(50)"`
	IconPath string `json:"icon_path" icon_path:"icon_path" gorm:"type:varchar(200)"`
}
