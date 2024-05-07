package utils

import "github.com/FourWD/middleware/common"

func NotiAcceptDeposit(notiToken string, userID string, depositID string) error {
	return common.SendMessageToUser(notiToken, "Accept Deposit", "Accept Deposit", nil)
}
