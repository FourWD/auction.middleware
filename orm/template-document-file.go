package orm

import "github.com/FourWD/middleware/model"

type TemplateDocumentFile struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	Name     string  `json:"name" query:"name" gorm:"type:varchar(256)"`
	RowOrder float32 `json:"row_order" query:"row_order" gorm:"type:decimal(5,2)"`
	IsActive bool    `json:"is_active" query:"is_active" gorm:"type:bool"`
}
