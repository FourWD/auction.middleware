package utils

import (
	"github.com/FourWD/middleware/common"
)

func NotiRemoveUserAll(notiToken string, userID string) error {
	topic := "ALL" // ให้ลบ user ออกจา่ก table NotiTopicUser
	return common.RemoveUserFromSubscription(topic, userID, notiToken)
}
