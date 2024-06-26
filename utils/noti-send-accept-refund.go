package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/FourWD/middleware/common"
	midOrm "github.com/FourWD/middleware/orm"
	"github.com/google/uuid"
)

func NotiSendAcceptRefund(userID string, refundID string) error {
	type Refund struct {
		Amount        int    `json:"amount"`
		ApproveDate   string `json:"approve_date"`
		BankAccountNo string `json:"bank_account_no"`
		BankName      string `json:"bank_name"`
	}
	refund := new(Refund)
	sqlRefund := `SELECT r.amount, DATE(r.approve_date) approve_date, r.bank_account_no, b.name_en bank_name FROM refunds r LEFT JOIN banks b ON r.bank_id = b.id WHERE r.id = ? AND r.user_id = ?`
	common.Database.Raw(sqlRefund, refundID, userID).Scan(refund)

	var notificationToken string
	sqlNotiToken := `SELECT notification_token FROM log_user_logins WHERE user_id = ? ORDER BY updated_at DESC LIMIT 1`
	common.Database.Raw(sqlNotiToken, userID).Scan(&notificationToken)

	if notificationToken == "" {
		common.PrintError("notificationToken", "notiToken is empty")
		return errors.New("notification token is empty")
	}

	obfuscated := refund.BankAccountNo
	if len(refund.BankAccountNo) > 3 {
		obfuscated = strings.Repeat("x", len(refund.BankAccountNo)-3) + refund.BankAccountNo[len(refund.BankAccountNo)-3:]
	}

	title := "🙏 Omakase Car Auction ขอขอบคุณ"
	body := fmt.Sprintf("คืนเงินมัดจําประมูล %d บาท %s เลขบัญชี %s %s", refund.Amount, refund.ApproveDate, obfuscated, refund.BankName)

	data := map[string]string{
		"user_id": userID,
	}

	if errSendMsg := common.SendMessageToUser(notificationToken, title, body, data); errSendMsg != nil {
		return errSendMsg
	}

	notification := midOrm.Notification{
		ID:                 uuid.NewString(),
		ToUserID:           userID,
		NotificationTypeID: "01",
		Message:            body,
		Title:              title,
		Url:                "",
		ShowDate:           time.Now(),
	}

	if err := common.Database.Debug().Create(&notification).Error; err != nil {
		return fmt.Errorf("failed to insert notification: %v", err)
	}

	return nil
}
