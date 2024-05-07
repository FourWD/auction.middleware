package utils

import "github.com/FourWD/middleware/common"

func NotiAcceptRefund(notiToken string, userID string, refundID string) error {
	return common.SendMessageToUser(notiToken, "Accept Refund", "Accept Refund", nil)
}
