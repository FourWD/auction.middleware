package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type BankTransfer struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID        string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
	UserID           string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	TitleID          string    `json:"title_id" query:"title_id" gorm:"type:varchar(20);"`
	PrefixID         string    `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2);"` //จาก int เป็น string
	Firstname        string    `json:"firstname" query:"firstname" gorm:"type:varchar(100);"`
	Lastname         string    `json:"lastname" query:"lastname" gorm:"type:varchar(100);"`
	IdcardNo         string    `json:"idcard_no" query:"idcard_no" gorm:"type:varchar(13)"`
	Telephone        string    `json:"telephone" query:"telephone" gorm:"type:varchar(30)"`
	FileIdcardID     string    `json:"file_idcard_id" query:"file_idcard_id" gorm:"type:varchar(36)"`
	TransferDateTime time.Time `json:"transfer_date_time" query:"transfer_date_time"`
	Amount           int       `json:"amount" query:"amount" gorm:"type:int"`
	FileSlipID       string    `json:"file_slip_id" query:"file_slip_id" gorm:"type:varchar(36)"`
	AcceptTCVersion  string    `json:"accept_tc_version" query:"accept_tc_version" gorm:"type:varchar(36)"`
	BankTransferID   string    `json:"bank_transfer_id" query:"bank_transfer_id" gorm:"type:varchar(2)"`
	ApproveBy        string    `json:"approve_by" query:"approve_by" gorm:"type:varchar(36)"`
	ApproveDate      time.Time `json:"approve_date" query:"approve_date"`
	IsMobileBanking  bool      `json:"is_mobile_banking" query:"is_mobile_banking" gorm:"type:bool"`
	OcupationID      string    `json:"ocupation_id" query:"ocupation_id" gorm:"type:varchar(36)"`

	Address       string `json:"address" query:"address" gorm:"type:text"`
	Street        string `json:"street" query:"street" gorm:"type:varchar(200)"`
	DistrictID    string `json:"district_id" query:"district_id" gorm:"type:varchar(4)"`         //อำเภอ
	SubDistrictID string `json:"sub_district_id" query:"sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ProvinceID    string `json:"province_id" query:"province_id" gorm:"type:varchar(2)"`

	FileIdCompanyRegister string `json:"file_id_company_register" query:"file_id_company_register" gorm:"type:varchar(36)"`
	FileIdPP20            string `json:"file_id_pp20" query:"file_id_pp20" gorm:"type:varchar(36)"`

	Remark   string `json:"remark" query:"remark" gorm:"type:text;"`
	Postcode string `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`
}
