package orm

import "github.com/FourWD/middleware/orm"

type RegisterAuctionMoneyTransfer struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	TransferDateTime     string `json:"article_type_id" query:"article_type_id" db:"article_type_id" gorm:"type:varchar(36)"`
	FileSlipID           string `json:"file_slip_id" query:"file_slip_id" db:"file_slip_id" gorm:"type:varchar(36)"`
	RefundBankID         string `json:"refund_bank_id" query:"refund_bank_id" db:"refund_bank_id" gorm:"type:varchar(36)"`
	RefundBankBranchName string `json:"refund_bank_branch_name" query:"refund_bank_branch_name" db:"refund_bank_branch_name" gorm:"type:varchar(200)"`
	RefundBankAccountNo  string `json:"refund_bank_account_no" query:"refund_bank_account_no" db:"refund_bank_account_no" gorm:"type:varchar(15)"`
	AcceptTCVersion      string `json:"accept_tc_version" query:"accept_tc_version" db:"accept_tc_version" gorm:"type:varchar(36)"`
	IsApproved           bool   `json:"is_approved" query:"is_approved" db:"is_approved" gorm:"type:boolean"`
	ApprovedBy           string `json:"approved_by" query:"approved_by" db:"approved_by" gorm:"type:varchar(36)"`
	ApprovedAt           string `json:"approved_at" query:"approved_at" db:"approved_at" gorm:"type:varchar(36)"`
}
