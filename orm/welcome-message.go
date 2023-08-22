package orm

import "github.com/FourWD/middleware/orm"

type WelcomeMessage struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	StartDate       string `json:"start_date" query:"start_date" db:"start_date" gorm:"not null;type:varchar(15)"`
	EndDate         string `json:"end_date" query:"end_date" db:"end_date" gorm:"not null;type:varchar(15)"`
	IsShow          bool   `json:"is_show" query:"is_show" db:"is_show" gorm:"bool"`
	ArticleID       string `json:"article_id" query:"article_id" db:"article_id" gorm:"type:varchar(36)"`
	CustomImagePath string `json:"custom_image_path" query:"custom_image_path" db:"custom_image_path" gorm:"type:varchar(50)"`
}
