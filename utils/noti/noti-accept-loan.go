package utils

import "github.com/FourWD/middleware/common"

func NotiAcceptLoan(userID string, loanID string) error {
	return common.SendMessageToUser(userID, "Accept Loan", "Accept Loan", nil)
}
