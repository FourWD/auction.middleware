package utils

import (
	"time"

	"github.com/FourWD/auction.middleware/orm"

	"github.com/FourWD/middleware/common"
	"github.com/google/uuid"
)

func SaveLogAuctionVehicleStatus(auctionID string, vehicleID string, auctionStatusID string) error {
	log := new(orm.LogAuctionVehicleStatus)
	log.ID = uuid.NewString()
	log.AuctionID = auctionID
	log.VehicleID = vehicleID
	log.AuctionStatusID = auctionStatusID
	log.CreatedAt = time.Now()

	if err := common.Database.Model(&orm.LogAuctionVehicleStatus{}).Create(&log); err.Error != nil {
		return err.Error
	}

	return nil

}
