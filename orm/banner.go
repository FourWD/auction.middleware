package orm

import "github.com/FourWD/middleware/orm"

type Banner struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BannerGroupID string `json:"banner_grounp_id" query:"banner_grounp_id" gorm:"type:varchar(36)"`

	Subject     string `json:"subject" query:"subject" gorm:"type:varchar(500)"`
	ImagePath   string `json:"image_path" query:"image_path" gorm:"type:varchar(100)"`
	Url         string `json:"url" query:"url" gorm:"type:varchar(200)"`
	YoutubeUrl  string `json:"youtube_url" query:"youtube_url" gorm:"type:varchar(100)"`
	Description string `json:"description" query:"description" gorm:"type:varchar(500)"`
	OpenType    string `json:"open_type" query:"open_type" gorm:"type:varchar(11)"`
	orm.RowOrder
}
