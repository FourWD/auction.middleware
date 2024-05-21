package utils

import (
	"errors"
	"fmt"

	midOrm "github.com/FourWD/middleware/orm"
	"github.com/google/uuid"

	"github.com/FourWD/middleware/common"
)

func NotiSendAcceptDeposit(userID string, depositID string) error {
	deposit := 0
	sqlDeposit := `SELECT amount FROM bank_transfers WHERE id = ? AND user_id = ?`
	common.Database.Raw(sqlDeposit, depositID, userID).Scan(&deposit)

	notificationToken := ""
	sqlNotiToken := `SELECT notification_token FROM log_user_logins WHERE user_id = ? ORDER BY updated_at DESC LIMIT 1`
	common.Database.Raw(sqlNotiToken, userID).Debug().Scan(&notificationToken)

	if notificationToken == "" {
		common.PrintError("notificationToken", "notiToken is empty")
		return errors.New("notiToken is empty")
	}

	title := "👏 ยินดีด้วย คุณได้รับสิทธิ์เข้าร่วมประมูล"
	body := fmt.Sprintf("ท่านมีสิทธิ์ชนะการประมูลได้ %d คัน", deposit/1000)

	data := map[string]string{
		"ีuser_id":   userID,
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
	}

	if err := common.Database.Debug().Create(&notification).Error; err != nil {
		return fmt.Errorf("failed to insert notification: %v", err)
	}

	return nil
}
