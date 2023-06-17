package orm

import "github.com/FourWD/middleware/orm"

type Banner struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BannerGroupID string `json:"banner_grounp_id" query:"banner_grounp_id" db:"banner_grounp_id" gorm:"type:varchar(36)"`
	
	ImagePath     string `json:"image_path" query:"image_path" db:"image_path" gorm:"type:varchar(100)"`
	Url           string `json:"url" query:"url" db:"url" gorm:"type:varchar(100)"`
	Youtube       string `json:"youtube" query:"youtube" db:"youtube" gorm:"type:varchar(100)"`
	Description   string `json:"description" query:"description" db:"description" gorm:"type:varchar(100)"`
}
