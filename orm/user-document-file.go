package orm

import "github.com/FourWD/middleware/model"

type UserDocumentFile struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`

	model.GormModel

	UserID                 string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	TemplateDocumentFileID string `json:"template_document_file_id" query:"template_document_file_id" gorm:"type:varchar(36)"`
	ImageID                string `json:"image_id" query:"image_id" gorm:"type:varchar(36)"`
	ImagePath              string `json:"image_path" query:"image_path" gorm:"type:varchar(256)"`
	IsDelete               bool   `json:"is_delete" query:"is_delete" gorm:"type:bool"`

	// FileIdcardID           string `json:"file_idcard_id" query:"file_idcard_id" gorm:"type:varchar(36)"`
	// FileCompanyRegisterID  string `json:"file_company_register_id" query:"file_company_register_id" gorm:"type:varchar(36)"`
	// FilePP20ID             string `json:"file_pp20_id" query:"file_pp20_id" gorm:"type:varchar(36)"`
	// FileHouseParticularsID string `json:"file_house_particulars_id" query:"file_house_particulars_id" gorm:"type:varchar(36)"`
}
