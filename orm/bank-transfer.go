package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type BankTransfer struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	UserID string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	// TitleID         string    `json:"title_id" query:"title_id" gorm:"type:varchar(20)"` //
	PrefixID             string    `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2)"` //จาก int เป็น string
	Firstname            string    `json:"firstname" query:"firstname" gorm:"type:varchar(100)"`
	Lastname             string    `json:"lastname" query:"lastname" gorm:"type:varchar(100)"`
	IdcardNo             string    `json:"idcard_no" query:"idcard_no" gorm:"type:varchar(13)"`
	Telephone            string    `json:"telephone" query:"telephone" gorm:"type:varchar(30)"`
	FileIdcardID         string    `json:"file_idcard_id" query:"file_idcard_id" gorm:"type:varchar(36)"`
	TransferDate         time.Time `json:"transfer_date" query:"transfer_date"`
	Amount               int       `json:"amount" query:"amount" gorm:"type:int"`
	FileSlipID           string    `json:"file_slip_id" query:"file_slip_id" gorm:"type:varchar(36)"`
	AcceptTCVersion      string    `json:"accept_tc_version" query:"accept_tc_version" gorm:"type:varchar(36)"`
	BankTransferStatusID string    `json:"bank_transfer_status_id" query:"bank_transfer_status_id" gorm:"type:varchar(2)"`
	ApproveBy            string    `json:"approve_by" query:"approve_by" gorm:"type:varchar(36)"`
	ApproveDate          time.Time `json:"approve_date" query:"approve_date"`
	IsMobileBanking      bool      `json:"is_mobile_banking" query:"is_mobile_banking" gorm:"type:bool"`
	IsApp                bool      `json:"is_app" query:"is_app" gorm:"type:bool"` //มาจากหน้าบ้าน = true , มาจากหลังบ้าน = false

	// PaymentTypeID string `json:"payment_type_id" query:"payment_type_id" gorm:"type:varchar(2)"`
	OccupationID string `json:"occupation_id" query:"occupation_id" gorm:"type:varchar(36)"`
	// EffectiveDate   time.Time `json:"effective_date" query:"effective_date"`

	Address       string `json:"address" query:"address" gorm:"type:text"`
	Street        string `json:"street" query:"street" gorm:"type:varchar(200)"`
	DistrictID    string `json:"district_id" query:"district_id" gorm:"type:varchar(4)"`         //อำเภอ
	SubDistrictID string `json:"sub_district_id" query:"sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ProvinceID    string `json:"province_id" query:"province_id" gorm:"type:varchar(2)"`

	FileCompanyRegisterID string `json:"file_company_register_id" query:"file_company_register_id" gorm:"type:varchar(36)"`
	FilePP20ID            string `json:"file_pp20_id" query:"file_pp20_id" gorm:"type:varchar(36)"`
	Code                  string `json:"code" query:"code" gorm:"type:varchar(20)"`
	RunningNo             int    `json:"running_no" query:"running_no" gorm:"type:int"`

	Remark   string `json:"remark" query:"remark" gorm:"type:text"`
	PostCode string `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`
}
