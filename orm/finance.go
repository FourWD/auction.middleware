package orm

import "github.com/FourWD/middleware/orm"

type Finance struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(10);primary_key;"`
	orm.GormModel

	Logo             string `json:"logo" query:"logo" gorm:"type:varchar(255)"`
	LogoHome         string `json:"logo_home" query:"logo_home" gorm:"type:varchar(255)"`
	Code             string `json:"code" query:"code" gorm:"type:varchar(4)"`
	Name             string `json:"name" query:"name" gorm:"type:varchar(50)"`
	Color            string `json:"color" query:"color" gorm:"type:varchar(7)"`
	LoanDiscription  string `json:"loan_discription" query:"loan_discription" gorm:"type:text"`
	Detail           string `json:"detail" query:"detail" gorm:"type:text"`
	Remark           string `json:"remark" query:"remark" gorm:"type:text"`
	IsRegisterActive bool   `json:"is_register_active" query:"is_register_active" gorm:"type:bool"`

	orm.GormRowOrder
}
