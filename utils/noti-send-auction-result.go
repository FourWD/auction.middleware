package utils

import (
	"fmt"
	"time"

	midOrm "github.com/FourWD/middleware/orm"
	"github.com/google/uuid"

	"github.com/FourWD/middleware/common"
)

func NotiSendAuctionResult(auctionID string) error {
	title := "ประกาศผล"
	body := "กดที่นี่เพื่อดูผลประมูล"

	data := map[string]string{
		"auction_id": auctionID,
		"event_code": "R0001",
	}
	topics := fmt.Sprintf(`JOIN_AUCTION_%s`, auctionID)

	if errSendToSubscriber := common.SendMessageToSubscriber(topics, title, body, data); errSendToSubscriber != nil {
		common.PrintError("errSendToSubscriber", errSendToSubscriber.Error())
		return errSendToSubscriber
	}
	topicid := ""
	sql := `SELECT id FROM notification_topics WHERE name = ?`
	common.Database.Raw(sql, topics).Scan(&topicid)
	notificationresult := midOrm.Notification{
		ID:                    uuid.NewString(),
		ToNotificationTopicID: topicid,
		NotificationTypeID:    "02",
		Message:               fmt.Sprintf(`"%s", "%s"`, title, body),
		ShowDate:              time.Now(),
	}

	if err := common.Database.Debug().Create(&notificationresult).Error; err != nil {
		return fmt.Errorf("failed to insert notificationresult: %v", err)
	}
	// type userJoinList struct {
	// 	UserID string `json:"user_id"`
	// }
	// var toUser []userJoinList
	// common.Database.Raw(`SELECT user_id FROM auction_vehicle_users WHERE auction_id = ? AND is_join = 1 GROUP BY user_id`, auctionID).Scan(&toUser)

	// for _, data := range toUser {
	// 	common.Database.Model(orm.Notification{}).Debug().Create(orm.Notification{
	// 		ID:                 uuid.NewString(),
	// 		ToUserID:           data.UserID,
	// 		NotificationTypeID: "01",
	// 		Message:            `{ tile : ` + title + ` body : ` + body + `}`,
	// 		Url:                "",
	// 	})
	// }
	return nil
}
