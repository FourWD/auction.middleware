package orm

import "github.com/FourWD/middleware/orm"

type File struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name      string `json:"article_type_id" query:"article_type_id" db:"article_type_id" gorm:"type:varchar(36)"`
	Extention string `json:"extention" query:"extention" db:"extention" gorm:"type:varchar(3)"`
	Path      string `json:"path" query:"path" db:"path" gorm:"type:varchar(100)"`
}
