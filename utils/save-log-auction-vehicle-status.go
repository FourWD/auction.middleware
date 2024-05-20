package utils

import (
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
	"github.com/google/uuid"
)

func LogAuctionVehicleStatus(auctionID string, vehicleID string, auctionStatus string) error {

	logAuctionVehicleStatus := new(orm.LogAuctionVehicleStatus)

	logAuctionVehicleStatus.ID = uuid.NewString()
	logAuctionVehicleStatus.AuctionID = auctionID
	logAuctionVehicleStatus.VehicleID = vehicleID
	logAuctionVehicleStatus.AuctionStatusID = auctionStatus
	logAuctionVehicleStatus.CreatedAt = time.Now()

	if err := common.Database.Model(&orm.LogAuctionStatus{}).Create(&logAuctionVehicleStatus); err.Error != nil {
		return err.Error
	}

	return nil

}
