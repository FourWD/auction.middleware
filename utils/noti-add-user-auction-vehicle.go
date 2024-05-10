package utils

import (
	"fmt"

	"github.com/FourWD/middleware/common"
)

func NotiAddUserAuctionVehicle(notiToken string, auctionID string, vehicleID string) error {
	fmt.Println(notiToken)
	topicAuction := fmt.Sprintf("JOIN_AUCTION_%s", auctionID)
	fmt.Println(topicAuction)
	if err := common.AddUserToSubscription(notiToken, topicAuction); err != nil {
		return err
	}

	topicAuctionVechile := fmt.Sprintf("JOIN_AUCTION_VEHICLE_%s_%s", auctionID, vehicleID)
	fmt.Println(topicAuctionVechile)
	return common.AddUserToSubscription(notiToken, topicAuctionVechile)
}
