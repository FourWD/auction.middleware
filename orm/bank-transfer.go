package orm

import "github.com/FourWD/middleware/orm"

type BankTransfer struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`

	TitleID string `json:"title_id" query:"title_id" gorm:"type:varchar(20);"`

	PrefixID     int    `json:"prefix_id" query:"prefix_id" gorm:"type:int(1);"`
	Firstname    string `json:"firstname" query:"firstname" gorm:"type:varchar(100);"`
	Lastname     string `json:"lastname" query:"lastname" gorm:"type:varchar(100);"`
	IdcardNO     string `json:"idcard_no" query:"idcard_no" gorm:"type:varchar(13)"`
	FileIdcardID string `json:"file_idcard_id" query:"file_idcard_id" gorm:"type:varchar(36)"`

	// โอนเงิน 3-2 เพิ่ใฟีลหน้าอัพโหลด
	TransferDateTime string `json:"article_type_id" query:"article_type_id" gorm:"type:varchar(36)"`
	FileSlipID       string `json:"file_slip_id" query:"file_slip_id" gorm:"type:varchar(36)"`
	AcceptTCVersion  string `json:"accept_tc_version" query:"accept_tc_version" gorm:"type:varchar(36)"`
	IsApproved       bool   `json:"is_approved" query:"is_approved" gorm:"type:bool"`
	ApprovedBy       string `json:"approved_by" query:"approved_by" gorm:"type:varchar(36)"`
	ApprovedAt       string `json:"approved_at" query:"approved_at" gorm:"type:varchar(36)"`
	IsMobileBanking  bool   `json:"is_mobile_banking" query:"is_mobile_banking" gorm:"type:bool"`
}
