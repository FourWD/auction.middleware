package orm

import "github.com/FourWD/middleware/orm"

type SubDistrict struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(6);primary_key;"`
	orm.GormModel

	DistrictID string `json:"district_id" query:"district_id" gorm:"type:varchar(4)"`
	Name       string `json:"name" query:"name" gorm:"not null;type:varchar(50)"`
	NameEn     string `json:"name_en" query:"name_en" gorm:"type:varchar(50)"`
	ZipCode    string `json:"zip_code" query:"zip_code" gorm:"type:varchar(5)"`
	orm.RowOrder
}
