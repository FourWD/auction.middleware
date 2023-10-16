package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type RegisterLeasing struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AuctionID              string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_register_leasings"`
	UserID                 string    `json:"user_id" query:"user_id" gorm:"type:varchar(36);uniqueIndex:idx_register_leasings"`
	PrefixID               int       `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2);"`
	Firstname              string    `json:"firstname" query:"firstname" gorm:"type:varchar(100);"`
	Lastname               string    `json:"lastname" query:"lastname" gorm:"type:varchar(100);"`
	IdcardNo               string    `json:"idcard_no" query:"idcard_no" gorm:"type:varchar(13)"`
	FileIDCardID           string    `json:"file_id_card_id" query:"file_id_card_id" gorm:"type:varchar(36)"`
	FileHouseParticularsID string    `json:"file_house_particulars_id" query:"file_house_particulars_id" gorm:"type:varchar(36)"`
	FilePayslipID          string    `json:"file_payslip_id" query:"file_payslip_id" gorm:"type:varchar(36)"`
	AcceptTCVersion        string    `json:"accept_tc_version" query:"accept_tc_version" gorm:"type:varchar(36)"`
	LeasingStatusID        string    `json:"leasing_status_id" query:"leasing_status_id" gorm:"type:varchar(36)"`
	FinancialAmount        string    `json:"financial_amount" query:"financial_amount" gorm:"type:varchar(36)"`
	UpdateLoanStatusBy     string    `json:"update_loan_status_by" query:"update_loan_status_by" gorm:"type:varchar(36)"`
	UpdateLoanStatusDate   time.Time `json:"update_loan_status_date" query:"update_loan_status_date" `
}
