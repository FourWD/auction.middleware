package orm

import "github.com/FourWD/middleware/orm"

type LogOTP struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	EmployeeID string `json:"employee_id" query:"employee_id" db:"employee_id" gorm:"type:varchar(36);"`
	UserID     string `json:"user_id" query:"user_id" db:"user_id" gorm:"type:varchar(36);"`
	OTP        string `json:"otp" query:"otp" db:"otp" gorm:"type:varchar(6);"`
	RefCodeOTP string `json:"ref_code_otp" query:"ref_code_otp" db:"ref_code_otp" gorm:"type:varchar(6);"`

	RequestDate string `json:"request_date" query:"request_date" db:"request_date" gorm:"type:varchar(15)"`
	SentDate    string `json:"sent_date" query:"sent_date" db:"sent_date" gorm:"type:varchar(15)"`
	ExpireDate  string `json:"expire_date" query:"expire_date" db:"expire_date" gorm:"type:varchar(15)"`
}
