package utils

import (
	"fmt"
	"strings"

	"github.com/FourWD/middleware/common"
	midOrm "github.com/FourWD/middleware/orm"
	"github.com/google/uuid"
)

func NotiAcceptRefund(userID string, refundID string) error {
	type Refund struct {
		Amount        int    `json:"amount"`
		ApproveDate   string `json:"approve_date"`
		BankAccountNo string `json:"bank_account_no"`
		BankName      string `json:"bank_name"`
	}
	refund := new(Refund)
	sqlRefund := `SELECT r.amount, DATE(r.approve_date) approve_date, r.bank_account_no, b.name_en bank_name FROM refunds r LEFT JOIN banks b ON r.bank_id = b.id WHERE id = ? AND user_id = ?`
	common.Database.Raw(sqlRefund, refundID, userID).Scan(&refund)

	notificationToken := ""
	sqlNotiToken := `SELECT notification_token FROM log_user_logins WHERE user_id = ? ORDER BY updated_at DESC LIMIT 1`
	common.Database.Raw(sqlNotiToken, userID).Debug().Scan(&notificationToken)

	obfuscated := strings.Repeat("x", len(refund.BankAccountNo)-3) + refund.BankAccountNo[len(refund.BankAccountNo)-3:]

	title := "🙏 Omakase Car Auction ขอขอบคุณ"
	body := fmt.Sprintf("คืนเงินมัดจําประมูล %d บาท %s เลขบัญชี %s %s", refund.Amount, refund.ApproveDate, obfuscated, refund.BankName)

	data := map[string]string{
		"ีuser_id":   userID,
		"event_code": "R0001",
	}

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
