package utils

import (
	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
	midOrm "github.com/FourWD/middleware/orm"
	"github.com/google/uuid"
)

func NotiAcceptDeposit(userID string, depositID string) error {
	deposit := new(orm.BankTransfer)
	sqlDeposit := `SELECT * FROM bank_transfers WHERE id = ? AND user_id = ?`
	common.Database.Raw(sqlDeposit, depositID, userID).Scan(&deposit)

	var notificationToken string
	sqlNotiToken := `SELECT notification_token FROM log_user_logins WHERE user_id = ?`
	common.Database.Raw(sqlNotiToken, userID).Debug().Scan(&notificationToken)

	title := "คำขอโอนเงินได้รับอนุมัติแล้ว"
	body := "กดที่นี่เพื่อดูผลคำขอโอนเงิน"

	data := map[string]string{
		"ีuser_id":   userID,
		"event_code": "R0001",
	}

	notificationToken = `fIW8kFalzU01uf2vI7ljNI:APA91bHhnQ9rOmCOWUcSSPhyo_LrrKaF58T3dhDkMFoo_M1xFGkFQL1RAhuYPHtugSFAQ9jBWUnxgJY3AZerBnqlgkQRsv6cpGhvjUC52A3qiRXjh8_X1E-Zug6sl6QAKQL6H8nM_X5g`

	if errSendMsg := common.SendMessageToUser(notificationToken, title, body, data); errSendMsg != nil {
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
