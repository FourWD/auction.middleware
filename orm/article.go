package orm

import "github.com/FourWD/middleware/orm"

type Article struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	ArticleTypeID string `json:"article_type_id" query:"article_type_id" db:"article_type_id" gorm:"type:varchar(36)"`
	FileCoverID   string `json:"file_cover_id" query:"file_cover_id" db:"file_cover_id" gorm:"type:varchar(36)"`
	Subject       string `json:"subject" query:"subject" db:"subject" gorm:"type:varchar(500)"`
	Detail        string `json:"detail" query:"detail" db:"detail" gorm:"type:varchar(500)"`
	ViewCount     int    `json:"view_count" query:"view_count" db:"view_count" gorm:"type:int(11)"`
	orm.RowOrder
}
