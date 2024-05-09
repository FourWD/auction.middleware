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

	logUserLogin := new(midOrm.LogUserLogin)
	sqlNotiToken := `SELECT * FROM log_user_logins WHERE user_id = ?`
	common.Database.Raw(sqlNotiToken, userID).Scan(&logUserLogin)

	title := "คำขอโอนเงินได้รับอนุมัติแล้ว"
	body := "กดที่นี่เพื่อดูผลคำขอโอนเงิน"

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
