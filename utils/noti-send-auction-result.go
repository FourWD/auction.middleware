package utils

import (
	"fmt"

	"github.com/FourWD/middleware/common"
	"github.com/FourWD/middleware/orm"
	"github.com/google/uuid"
)

func NotiSendAuctionResult(auctionID string) error {
	title := "ประกาศผล"
	body := "กดที่นี่เพื่อดูผลประมูล"

	data := map[string]string{
		"auction_id": auctionID,
		"event_code": "R0001",
	}
	common.SendMessageToSubscriber(fmt.Sprintf(`JOIN_AUCTION_%s`, auctionID), title, body, data)

	type userJoinList struct {
		UserID string `json:"user_id"`
	}
	var toUser []userJoinList
	common.Database.Raw(`SELECT user_id FROM auction_vehicle_users WHERE auction_id = '` + auctionID + `' AND is_join = 1 GROUP BY user_id `).Scan(&toUser)

	for _, data := range toUser {
		common.Database.Model(orm.Notification{}).Create(orm.Notification{
			ID:                 uuid.NewString(),
			ToUserID:           data.UserID,
			NotificationTypeID: "01",
			Message:            `{ tile : ` + title + ` body : ` + body + `}`,
			Url:                "",
		})
	}
	return nil
}
