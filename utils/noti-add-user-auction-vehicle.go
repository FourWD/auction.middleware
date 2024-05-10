package utils

import (
	"errors"
	"fmt"

	"github.com/FourWD/middleware/common"
)

func NotiAddUserAuctionVehicle(notiToken string, auctionID string, vehicleID string) error {
	if notiToken == "" || auctionID == "" || vehicleID == "" {
		common.PrintError("NotiAddUserAuctionVehicle", "some value is empty notiToken : "+notiToken+" auctionID : "+auctionID+" vehicleID : "+vehicleID)
		return errors.New("some value is empty")
	}
	common.Print("notiToken", notiToken)
	topicAuction := fmt.Sprintf("JOIN_AUCTION_%s", auctionID)
	common.Print("topicAuction", topicAuction)
	if err := common.AddUserToSubscription(notiToken, topicAuction); err != nil {
		return err
	}

	topicAuctionVechile := fmt.Sprintf("JOIN_AUCTION_VEHICLE_%s_%s", auctionID, vehicleID)
	common.Print("topicAuctionVechile", topicAuctionVechile)
	return common.AddUserToSubscription(notiToken, topicAuctionVechile)
}
