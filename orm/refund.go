package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Refund struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BankID          string    `json:"refund_bank_id" query:"refund_bank_id" gorm:"type:varchar(36)"`
	BankBranchName  string    `json:"refund_bank_branch_name" query:"refund_bank_branch_name" gorm:"type:varchar(200)"`
	BankAccountName string    `json:"refund_bank_account_name" query:"refund_bank_account_name" gorm:"type:varchar(15)"`
	BankAccountNo   string    `json:"refund_bank_account_no" query:"refund_bank_account_no" gorm:"type:varchar(15)"`
	RightToBid      int       `json:"right_to_bid" query:"right_to_bid" gorm:"type:int(5)"`
	Amount          float64   `json:"amount" query:"amount" gorm:"type:float"`
	RequestDate     time.Time `json:"request_date" query:"request_date"`
	RefundStatusID  string    `json:"refund_status_id" query:"refund_status_id" gorm:"type:varchar(36)"`
	IsApprove       bool      `json:"is_approve" query:"is_approve" gorm:"type:bool"`
	ApproveBy       string    `json:"approve_by" query:"approve_by" gorm:"type:varchar(36)"`
	ApproveDate     time.Time `json:"approve_date" query:"approve_date"`
	Remark          string    `json:"remark" query:"remark" gorm:"type:text;"`
}
