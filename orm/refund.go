package orm

import "github.com/FourWD/middleware/orm"

type Refund struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BankID          string  `json:"refund_bank_id" query:"refund_bank_id" gorm:"type:varchar(36)"`
	BankBranchName  string  `json:"refund_bank_branch_name" query:"refund_bank_branch_name" gorm:"type:varchar(200)"`
	BankAccountName string  `json:"refund_bank_account_name" query:"refund_bank_account_name" gorm:"type:varchar(15)"`
	BankAccountNo   string  `json:"refund_bank_account_no" query:"refund_bank_account_no" gorm:"type:varchar(15)"`
	RightToBid      int     `json:"right_to_bid" query:"right_to_bid" gorm:"type:int(5)"`
	Amount          float64 `json:"amount" query:"amount" gorm:"type:float"`
	RequestDate     string  `json:"request_date" query:"request_date" gorm:"type:varchar(20)"`
	RefundStatusID  string  `json:"refund_status_id" query:"refund_status_id" gorm:"type:varchar(36)"`
	IsApproved      bool    `json:"is_approved" query:"is_approved" gorm:"type:bool"`
	ApprovedBy      string  `json:"approved_by" query:"approved_by" gorm:"type:varchar(36)"`
	ApprovedAt      string  `json:"approved_at" query:"approved_at" gorm:"type:varchar(36)"`
	Remark          string  `json:"remark" query:"remark" db:"name" gorm:"type:text;"`
}
