package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	midOrm "github.com/FourWD/middleware/orm"
	"github.com/google/uuid"

	"github.com/FourWD/middleware/common"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func NotiSendAcceptLoan(userID string, loanID string) error {
	loan := ""
	sqlLoan := `SELECT financial_amount FROM register_leasings WHERE id = ? AND user_id = ?`
	common.Database.Raw(sqlLoan, loanID, userID).Scan(&loan)

	notificationToken := ""
	sqlNotiToken := `SELECT notification_token FROM log_user_logins WHERE user_id = ? ORDER BY updated_at DESC LIMIT 1`
	common.Database.Raw(sqlNotiToken, userID).Debug().Scan(&notificationToken)

	if notificationToken == "" {
		common.PrintError("notificationToken", "notiToken is empty")
		return errors.New("notiToken is empty")
	}

	loanInt, _ := strconv.Atoi(loan)
	p := message.NewPrinter(language.English)
	loanWithCommaThousandSep := p.Sprintf("%d", loanInt)
	title := "👏 ยินดีด้วย คุณได้รับสิทธิ์เข้าร่วมประมูล"
	body := fmt.Sprintf("วงเงินเช่าซื้อเบื้อง %s บาท (เงื่อนไขเป็นไปตามที่สถาบันการเงินกําหนด)", loanWithCommaThousandSep)

	data := map[string]string{
		"user_id":    userID,
		"event_code": "R0001",
	}

	if errSendMsg := common.SendMessageToUser(notificationToken, title, body, data); errSendMsg != nil {
		return errSendMsg
	}

	notification := midOrm.Notification{
		ID:                 uuid.NewString(),
		ToUserID:           userID,
		NotificationTypeID: "01",
		Message:            fmt.Sprintf(`{"title": "%s", "body": "%s"}`, title, body),
		Url:                "",
		ShowDate:           time.Now(),
	}

	if err := common.Database.Debug().Create(&notification).Error; err != nil {
		return fmt.Errorf("failed to insert notification: %v", err)
	}

	return nil
}
