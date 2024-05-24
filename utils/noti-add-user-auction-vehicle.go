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
	topicAuctionVehicle := fmt.Sprintf("JOIN_AUCTION_VEHICLE_%s_%s", auctionID, vehicleID)
	common.Print("topicAuctionVechile", topicAuctionVehicle)

	// notificationtopics := midOrm.NotificationTopic{
	// 	ID:   uuid.NewString(),
	// 	Name: topicAuctionVechile,
	// }

	// if err := common.Database.Debug().Create(&notificationtopics).Error; err != nil {
	// 	return fmt.Errorf("failed to insert notificationtopics: %v", err)
	// }
	// var existingTopic midOrm.NotificationTopic
	// err := common.Database.Where("name = ?", topicAuctionVehicle).First(&existingTopic).Error
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		notificationTopic := midOrm.NotificationTopic{
	// 			ID:   uuid.NewString(),
	// 			Name: topicAuctionVehicle,
	// 		}
	// 		if err := common.Database.Debug().Create(&notificationTopic).Error; err != nil {
	// 			common.PrintError("Failed to insert notification topic: ", err.Error())
	// 		}
	// 	} else {
	// 		common.PrintError("Failed to check existing topic: ", err.Error())
	// 	}
	// } else {
	// 	common.Print("Topic already exists: ", topicAuctionVehicle)
	// }
	return common.AddUserToSubscription(notiToken, topicAuctionVehicle)
}

func NotiRemoveUserAuctionVehicle(notiToken string, auctionID string, vehicleID string) error {
	if notiToken == "" || auctionID == "" || vehicleID == "" {
		common.PrintError("NotiRemoveUserAuctionVehicle", "some value is empty notiToken : "+notiToken+" auctionID : "+auctionID+" vehicleID : "+vehicleID)
		return errors.New("some value is empty")
	}
	common.Print("notiToken", notiToken)
	topicAuction := fmt.Sprintf("JOIN_AUCTION_%s", auctionID)
	common.Print("topicAuction", topicAuction)
	if err := common.RemoveUserFromSubscription(notiToken, topicAuction); err != nil {
		return err
	}

	topicAuctionVechile := fmt.Sprintf("JOIN_AUCTION_VEHICLE_%s_%s", auctionID, vehicleID)
	common.Print("topicAuctionVechile", topicAuctionVechile)
	return common.RemoveUserFromSubscription(notiToken, topicAuctionVechile)
}
