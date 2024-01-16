package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type RegisterLeasing struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id"`
	model.GormModel

	RequestDate            time.Time `json:"request_date" query:"request_date"`
	Code                   string    `json:"code" query:"code" gorm:"type:varchar(20)"` //รหัสยื่นเรื่อง CS / FN
	UserID                 string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	FinanceID              string    `json:"finance_id" query:"finance_id" gorm:"type:varchar(10)"`
	PrefixID               string    `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2)"`
	Firstname              string    `json:"firstname" query:"firstname" gorm:"type:varchar(100)"`
	Lastname               string    `json:"lastname" query:"lastname" gorm:"type:varchar(100)"`
	IdcardNo               string    `json:"idcard_no" query:"idcard_no" gorm:"type:varchar(13)"`
	FileIDCardID           string    `json:"file_id_card_id" query:"file_id_card_id" gorm:"type:varchar(36)"`
	FileHouseParticularsID string    `json:"file_house_particulars_id" query:"file_house_particulars_id" gorm:"type:varchar(36)"`
	FilePayslipID          string    `json:"file_payslip_id" query:"file_payslip_id" gorm:"type:varchar(36)"`
	Telephone              string    `json:"telephone" query:"telephone" gorm:"type:varchar(30)"`

	Address       string `json:"address" query:"address" gorm:"type:text"`
	Street        string `json:"street" query:"street" gorm:"type:varchar(200)"`
	DistrictID    string `json:"district_id" query:"district_id" gorm:"type:varchar(4)"`         //อำเภอ
	SubDistrictID string `json:"sub_district_id" query:"sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ProvinceID    string `json:"province_id" query:"province_id" gorm:"type:varchar(2)"`

	OcupationID      string    `json:"ocupation_id" query:"ocupation_id" gorm:"type:varchar(36)"`
	NetIncome        int       `json:"net_income" query:"net_income" gorm:"type:int"`
	IsLoanSCB        bool      `json:"is_loan_scb" query:"is_loan_scb" gorm:"type:bool"`                 //มีสินเชื่อรถกับSCB ?
	IsBankAccountSCB bool      `json:"is_bank_account_scb" query:"is_bank_account_scb" gorm:"type:bool"` //มีบัญชีSCB ?
	IsApp            bool      `json:"is_app" query:"is_app" gorm:"type:bool"`                           //มาจากหน้าบ้าน = true , มาจากหลังบ้าน = false
	BirthDate        time.Time `json:"birth_date" query:"birth_date"`
	YearExperience   int       `json:"year_experience" query:"year_experience" gorm:"type:int"` //ปีที่ทำงาน

	AcceptTCVersion      string    `json:"accept_tc_version" query:"accept_tc_version" gorm:"type:varchar(36)"`
	LeasingStatusID      string    `json:"leasing_status_id" query:"leasing_status_id" gorm:"type:varchar(36)"`
	FinanceStatusID      string    `json:"finance_status_id" query:"finance_status_id" gorm:"type:varchar(36)"`
	FinancialAmount      string    `json:"financial_amount" query:"financial_amount" gorm:"type:varchar(36)"`
	UpdateLoanStatusBy   string    `json:"update_loan_status_by" query:"update_loan_status_by" gorm:"type:varchar(36)"`
	UpdateLoanStatusDate time.Time `json:"update_loan_status_date" query:"update_loan_status_date" `
	IsExpire             bool      `json:"is_expire" query:"is_expire" gorm:"type:bool"`
	ExpireDate           time.Time `json:"expire_date" query:"expire_date"`

	LastOperationUpdateBy   string    `json:"last_operation_update_by" query:"last_operation_update_by" gorm:"type:varchar(36)"`
	LastOperationUpdateDate time.Time `json:"last_operation_update_date" query:"last_operation_update_date" `

	LastFinanceUpdateBy   string    `json:"last_finance_update_by" query:"last_finance_update_by" gorm:"type:varchar(36)"`
	LastFinanceUpdateDate time.Time `json:"last_finance_update_date" query:"last_finance_update_date" `
	RunningNo             int       `json:"running_no" query:"running_no" gorm:"type:int"`
}
