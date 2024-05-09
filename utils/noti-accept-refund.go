package utils

import (
	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
	midOrm "github.com/FourWD/middleware/orm"
	"github.com/google/uuid"
)

func NotiAcceptRefund(userID string, refundID string) error {
	refund := new(orm.Refund)
	sqlRefund := `SELECT * FROM refunds WHERE id = ? AND user_id = ?`
	common.Database.Raw(sqlRefund, refundID, userID).Scan(&refund)

	logUserLogin := new(midOrm.LogUserLogin)
	sqlNotiToken := `SELECT * FROM log_user_logins WHERE user_id = ?`
	common.Database.Raw(sqlNotiToken, userID).Scan(&logUserLogin)

	title := "คำขอคืนเงินได้รับอนุมัติแล้ว"
	body := "กดที่นี่เพื่อดูผลคำขอคืนเงิน"

	data := map[string]string{
		"ีuser_id":   userID,
		"event_code": "R0001",
	}

	if errSendMsg := common.SendMessageToUser(logUserLogin.NotificationToken, title, body, data); errSendMsg != nil {
		return errSendMsg
	}

	common.Database.Model(midOrm.Notification{}).Debug().Create(midOrm.Notification{
		ID:                 uuid.NewString(),
		ToUserID:           userID,
		NotificationTypeID: "01",
		Message:            `{ tile : ` + title + ` body : ` + body + `}`,
		Url:                "",
	})

	return nil
}
