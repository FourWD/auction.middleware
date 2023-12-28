package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Refund struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	// BankID          string    `json:"refund_bank_id" query:"refund_bank_id" gorm:"type:varchar(36)"`
	UserID          string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	BankBranchName  string    `json:"refund_bank_branch_name" query:"refund_bank_branch_name" gorm:"type:varchar(200)"`
	BankListID      string    `json:"bank_list_id" query:"bank_list_id" gorm:"type:varchar(2)"`
	BankAccountName string    `json:"refund_bank_account_name" query:"refund_bank_account_name" gorm:"type:varchar(15)"`
	BankAccountNo   string    `json:"refund_bank_account_no" query:"refund_bank_account_no" gorm:"type:varchar(15)"`
	RightToBid      int       `json:"right_to_bid" query:"right_to_bid" gorm:"type:int(5)"`
	Amount          int       `json:"amount" query:"amount" gorm:"type:int"`
	RequestDate     time.Time `json:"request_date" query:"request_date"`
	// EffectiveDate   time.Time `json:"effective_date" query:"effective_date"`
	RefundStatusID string    `json:"refund_status_id" query:"refund_status_id" gorm:"type:varchar(2)"`
	IsApprove      bool      `json:"is_approve" query:"is_approve" gorm:"type:bool"`
	ApproveBy      string    `json:"approve_by" query:"approve_by" gorm:"type:varchar(36)"`
	ApproveDate    time.Time `json:"approve_date" query:"approve_date"`
	IsTransfer     bool      `json:"is_transfer" query:"is_transfer" gorm:"type:bool"`
	TransferBy     string    `json:"transfer_by" query:"transfer_by" gorm:"type:varchar(36)"`
	TransfereDate  time.Time `json:"transfer_date" query:"transfer_date"`
	Remark         string    `json:"remark" query:"remark" gorm:"type:text"`
	FileSlipID     string    `json:"file_slip_id" query:"file_slip_id" gorm:"type:varchar(36)"`
	IsApp          bool      `json:"is_app" query:"is_app" gorm:"type:bool"` //มาจากหน้าบ้าน = true , มาจากหลังบ้าน = false

}
