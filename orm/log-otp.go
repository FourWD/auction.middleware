package orm

import (
	"github.com/FourWD/middleware/orm"
)

type LogOTP struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	EmployeeID string `json:"employee_id" query:"employee_id" gorm:"type:varchar(36);"`
	UserID     string `json:"user_id" query:"user_id" gorm:"type:varchar(36);"`
	DeviceID   string `json:"device_id" query:"device_id"  gorm:"type:varchar(255)"`
	OTP        string `json:"otp" query:"otp" gorm:"type:varchar(6);"`
	RefCodeOTP string `json:"ref_code_otp" query:"ref_code_otp" gorm:"type:varchar(6);"`

	RequestDate string `json:"request_date" query:"request_date" gorm:"type:time"`
	SentDate    string `json:"sent_date" query:"sent_date" gorm:"type:time"`
	VerifyDate  string `json:"verify_date" query:"verify_date" gorm:"type:time"`
	ExpireDate  string `json:"expire_date" query:"expire_date" gorm:"type:time"`
}
