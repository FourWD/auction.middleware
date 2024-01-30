package orm

import "github.com/FourWD/middleware/orm"

type UserProfile struct {
	orm.UserProfile
	CompanyVatTypeID      string `json:"company_vat_type_id" query:"company_vat_type_id" gorm:"type:varchar(2)"`
	CompanyRegisterTypeID string `json:"company_register_type_id" query:"company_register_type_id" gorm:"type:varchar(2)"`
	BusinessTypeID        string `json:"business_type_id" query:"business_type_id" gorm:"type:varchar(2)"`
	CompanyName           string `json:"company_name" query:"company_name" gorm:"type:varchar(50)"`
	CompanyPhone          string `json:"company_phone" query:"company_phone" gorm:"type:varchar(50)"`
	Tax                   string `json:"tax" query:"tax" gorm:"type:varchar(50)"`
	HQShowroomName        string `json:"hq_show_room_name" query:"hq_show_room_name" gorm:"type:varchar(500)"`
	BankID                string `json:"bank_id" query:"bank_id" gorm:"type:varchar(2)"`
	BankAccountName       string `json:"refund_bank_account_name" query:"refund_bank_account_name" gorm:"type:varchar(15)"`
	BankAccountNo         string `json:"refund_bank_account_no" query:"refund_bank_account_no" gorm:"type:varchar(15)"`

	FileIdcardID           string `json:"file_idcard_id" query:"file_idcard_id" gorm:"type:varchar(36)"`
	FileCompanyRegisterID  string `json:"file_company_register_id" query:"file_company_register_id" gorm:"type:varchar(36)"`
	FilePP20ID             string `json:"file_pp20_id" query:"file_pp20_id" gorm:"type:varchar(36)"`
	FileHouseParticularsID string `json:"file_house_particulars_id" query:"file_house_particulars_id" gorm:"type:varchar(36)"`
}
