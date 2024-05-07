package utils

import (
	"github.com/FourWD/middleware/common"
)

func NotiAddUserAll(notiToken string) error {
	topic := "ALL"
	return common.AddUserToSubscription(notiToken, topic)
}
