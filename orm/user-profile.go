package orm

import "github.com/FourWD/middleware/orm"

type UserProfile struct {
	orm.UserProfile

	FileAttorneyLetterID string `json:"file_attorney_letter_id" query:"file_attorney_letter_id" gorm:"type:varchar(36)"`
	FileAttorneyIdcardID string `json:"file_attorney_idcard_id" query:"file_attorney_idcard_id" gorm:"type:varchar(36)"`

	IsAddressTax bool `json:"is_address_tax" query:"is_address_tax" gorm:"type:bool"` //สำหรับ ติ๊ก
}
