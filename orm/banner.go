package orm

import "github.com/FourWD/middleware/orm"

type Banner struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BannerGroupID string `json:"banner_grounp_id" query:"banner_grounp_id" db:"banner_grounp_id" gorm:"type:varchar(36)"`

	Subject     string `json:"subject" query:"subject" db:"subject" gorm:"type:varchar(500)"`
	ImagePath   string `json:"image_path" query:"image_path" db:"image_path" gorm:"type:varchar(100)"`
	Url         string `json:"url" query:"url" db:"url" gorm:"type:varchar(200)"`
	YoutubeUrl  string `json:"youtube_url" query:"youtube_url" db:"youtube_url" gorm:"type:varchar(100)"`
	Description string `json:"description" query:"description" db:"description" gorm:"type:varchar(500)"`
	OpenType    string `json:"open_type" query:"open_type" db:"open_type" gorm:"type:varchar(11)"`
	orm.RowOrder
}
