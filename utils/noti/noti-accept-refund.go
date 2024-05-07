package utils

import "github.com/FourWD/middleware/common"

func NotiAcceptRefund(userID string, refundID string) error {
	return common.SendMessageToUser(userID, "Accept Refund", "Accept Refund", nil)
}
