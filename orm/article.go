package orm

import (
	"github.com/FourWD/middleware/orm"
)

type Article struct {
	orm.Article
	ButtonName  string `json:"button_name" query:"button_name" gorm:"type:varchar(100)"`
	ButtonUrl   string `json:"button_url" query:"button_url" gorm:"type:varchar(100)"`
	IsShowOnApp bool   `json:"is_show_on_app" query:"is_show_on_app" gorm:"type:bool"`
}
