package orm

import "github.com/FourWD/middleware/orm"

type UserProfile struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID                string `json:"user_id" query:"user_id" gorm:"type:varchar(36);"`
	CompanyVatTypeID      string `json:"company_vat_type_id" query:"company_vat_type_id" gorm:"type:varchar(2);"`
	CompanyRegisterTypeID string `json:"company_register_type_id" query:"company_register_type_id" gorm:"type:varchar(2);"`
	BusinessTypeID        string `json:"business_type_id" query:"business_type_id" gorm:"type:varchar(2);"`
	CompanyName           string `json:"company_name" query:"company_name" gorm:"type:varchar(50);"`
	CompanyPhone          string `json:"company_phone" query:"company_phone" gorm:"type:varchar(50);"`
	Tax                   string `json:"tax" query:"tax" gorm:"type:varchar(50);"`
	HQShowroomName        string `json:"hq_show_room_name" query:"hq_show_room_name" gorm:"type:varchar(500);"`
}
