package orm

import "github.com/FourWD/middleware/orm"

type NotificationType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name      string `json:"name" query:"name" gorm:"not null;type:varchar(50)"`
	ImagePath string `json:"image_path" query:"image_path" gorm:"type:varchar(200)"`
	orm.RowOrder
}
