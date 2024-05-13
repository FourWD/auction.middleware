package utils

import (
	"github.com/FourWD/middleware/common"
)

func NotiRemoveUserAll(notiToken string) error {
	topic := "ALL" // ให้ลบ user ออกจา่ก table NotiTopicUser
	return common.RemoveUserFromSubscription(notiToken, topic)
}
