package orm

import "github.com/FourWD/middleware/orm"

type RegisterAuctionLoan struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FileIDCardID           string `json:"file_id_card_id" query:"file_id_card_id" db:"file_id_card_id" gorm:"type:varchar(36)"`
	FileHouseParticularsID string `json:"file_house_particulars_id" query:"file_house_particulars_id" db:"file_house_particulars_id" gorm:"type:varchar(36)"`
	FilePayslipID          string `json:"file_payslip_id" query:"file_payslip_id" db:"file_payslip_id" gorm:"type:varchar(36)"`
	AcceptTCVersion        string `json:"accept_tc_version" query:"accept_tc_version" db:"accept_tc_version" gorm:"type:varchar(36)"`

	LoanStatusID        string `json:"loan_status_id" query:"loan_status_id" db:"loan_status_id" gorm:"type:varchar(36)"`
	FinancialAmount     string `json:"financial_amount" query:"financial_amount" db:"financial_amount" gorm:"type:varchar(36)"`
	UpdatedLoanStatusBy string `json:"updated_loan_status_by" query:"updated_loan_status_by" db:"updated_loan_status_by" gorm:"type:varchar(36)"`
	UpdatedLoanStatusAt string `json:"updated_loan_status_at" query:"updated_loan_status_at" db:"updated_loan_status_at" gorm:"type:varchar(36)"`
}
