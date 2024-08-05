package orm

import (
	"github.com/FourWD/middleware/model"
)

type RoundPayment struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id"`
	model.GormModel

	BankName      string `json:"bank_name" query:"bank_name" gorm:"type:varchar(255)"`
	BankAccount   string `json:"bank_account" query:"bank_account" gorm:"type:varchar(255)"`
	BankAccountNo string `json:"bank_account_no" query:"bank_account_no" gorm:"type:varchar(20)"`

	BankQr string `json:"bank_qr" query:"bank_qr" gorm:"type:varchar(255)"`
}
