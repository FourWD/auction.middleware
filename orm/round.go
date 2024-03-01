package orm

import (
	"github.com/FourWD/middleware/model"
)

type Round struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(10);primary_key"`
	model.GormModel

	Logo     string `json:"logo" query:"logo" gorm:"type:varchar(255)"`
	LogoHome string `json:"logo_home" query:"logo_home" gorm:"type:varchar(255)"`
	Code     string `json:"code" query:"code" gorm:"type:varchar(4)"`
	Name     string `json:"name" query:"name" gorm:"type:varchar(200)"`
	Color    string `json:"color" query:"color" gorm:"type:varchar(7)"`
	IsActive bool   `json:"is_active" query:"is_active" gorm:"type:bool"`

	model.GormRowOrder
}
