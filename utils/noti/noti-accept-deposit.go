package utils

import "github.com/FourWD/middleware/common"

func NotiAcceptDeposit(userID string, depositID string) error {
	return common.SendMessageToUser(userID, "Accept Deposit", "Accept Deposit", nil)
}
