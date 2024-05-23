package utils

import (
	"fmt"

	"github.com/FourWD/middleware/common"
)

func NotiAddUserAll(notiToken string) error {
	topic := "ALL"
	return common.AddUserToSubscription(notiToken, topic)
}

func AddUserToAuctionGroup(userToken string, auctionID string, vehicleID string) error {
	topic := fmt.Sprintf("auction_%s_vehicle_%s", auctionID, vehicleID)
	return common.AddUserToSubscription(userToken, topic)
}
