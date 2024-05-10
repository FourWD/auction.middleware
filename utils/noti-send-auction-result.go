package utils

import (
	"fmt"

	"github.com/FourWD/middleware/common"
)

func NotiSendAuctionResult(auctionID string) error {
	title := "ประกาศผล"
	body := "กดที่นี่เพื่อดูผลประมูล"

	data := map[string]string{
		"auction_id": auctionID,
		"event_code": "R0001",
	}
	if errSendToSubscriber := common.SendMessageToSubscriber(fmt.Sprintf(`JOIN_AUCTION_%s`, auctionID), title, body, data); errSendToSubscriber != nil {
		common.PrintError("errSendToSubscriber", errSendToSubscriber.Error())
		return errSendToSubscriber
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
