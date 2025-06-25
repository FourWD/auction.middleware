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
	// ========================================================================================
	logFields := map[string]interface{}{
		"auction_id":        auctionID,
		"vehicle_id":        vehicleID,
		"auction_status_id": auctionStatusID,
		"created_at":        log.CreatedAt,
	}
	// ========================================================================================
	if err := common.Database.Model(&orm.LogAuctionVehicleStatus{}).Create(&log); err.Error != nil {
		logFields["error"] = err.Error
		common.LogError("SaveLogAuctionVehicleStatus", logFields, vehicleID)
		return err.Error
	}
	// ========================================================================================
	common.Log("SaveLogAuctionVehicleStatus", logFields, vehicleID)
	return nil

}
