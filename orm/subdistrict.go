package orm

import "github.com/FourWD/middleware/orm"

type Subdistrict struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DistrictIDID string `json:"district_id" query:"district_id" gorm:"type:varchar(2)"`
	Name         string `json:"name" query:"name" gorm:"not null;type:varchar(100)"`
	NameEn       string `json:"name_en" query:"name_en" gorm:"type:varchar(50)"`
}
