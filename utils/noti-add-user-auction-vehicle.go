package utils

import (
	"errors"
	"fmt"

	"github.com/FourWD/middleware/common"
)

func NotiAddUserAuctionVehicle(notiToken string, userID string, auctionID string, vehicleID string) error {
	if notiToken == "" || auctionID == "" || vehicleID == "" {
		// common.PrintError("NotiAddUserAuctionVehicle", "some value is empty notiToken : "+notiToken+" auctionID : "+auctionID+" vehicleID : "+vehicleID)
		return errors.New("some value is empty")
	}
	// common.Print("notiToken", notiToken)
	topicAuction := fmt.Sprintf("JOIN_AUCTION_%s", auctionID)
	// common.Print("topicAuction", topicAuction)

	if err := common.AddUserToSubscription(topicAuction, userID, notiToken); err != nil {
		return err
	}
	topicAuctionVehicle := fmt.Sprintf("JOIN_AUCTION_VEHICLE_%s_%s", auctionID, vehicleID)
	// common.Print("topicAuctionVechile", topicAuctionVehicle)

	return common.AddUserToSubscription(topicAuctionVehicle, userID, notiToken)
}

func NotiRemoveUserAuctionVehicle(notiToken string, userID string, auctionID string, vehicleID string) error {
	if notiToken == "" || auctionID == "" || vehicleID == "" {
		// common.PrintError("NotiRemoveUserAuctionVehicle", "some value is empty notiToken : "+notiToken+" auctionID : "+auctionID+" vehicleID : "+vehicleID)
		return errors.New("some value is empty")
	}
	// common.Print("notiToken", notiToken)
	topicAuction := fmt.Sprintf("JOIN_AUCTION_%s", auctionID)
	// common.Print("topicAuction", topicAuction)
	if err := common.RemoveUserFromSubscription(topicAuction, userID, notiToken); err != nil {
		return err
	}

	topicAuctionVechile := fmt.Sprintf("JOIN_AUCTION_VEHICLE_%s_%s", auctionID, vehicleID)
	// common.Print("topicAuctionVechile", topicAuctionVechile)
	return common.RemoveUserFromSubscription(topicAuctionVechile, userID, notiToken)
}
