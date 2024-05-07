package utils

import "github.com/FourWD/middleware/common"

func NotiAcceptLoan(notiToken string, userID string, loanID string) error {
	return common.SendMessageToUser(notiToken, "Accept Loan", "Accept Loan", nil)
}
