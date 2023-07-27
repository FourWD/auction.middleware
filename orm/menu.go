package orm

import "github.com/FourWD/middleware/orm"

type Menu struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	MenuGroupID string `json:"menu_grounp_id" query:"menu_grounp_id" db:"menu_grounp_id" gorm:"type:varchar(36)"`

	Name     string `json:"name" query:"name" db:"name" gorm:"type:varchar(500)"`
	IconPath string `json:"icon_path" query:"icon_path" db:"icon_path" gorm:"type:varchar(100)"`
	Url      string `json:"url" query:"url" db:"url" gorm:"type:varchar(200)"`
	OpenType string `json:"open_type" query:"open_type" db:"open_type" gorm:"type:varchar(10)"`
	orm.RowOrder
}
