package utils

import (
	"time"

	"github.com/FourWD/auction.middleware/orm"

	"github.com/FourWD/middleware/common"
	"github.com/FourWD/middleware/infra"
	"github.com/google/uuid"
)

func SaveLogAuctionVehicleStatus(auctionID string, vehicleID string, auctionStatusID string) error {
	log := &orm.LogAuctionVehicleStatus{
		ID:              uuid.NewString(),
		AuctionID:       auctionID,
		VehicleID:       vehicleID,
		AuctionStatusID: auctionStatusID,
	}
	log.CreatedAt = time.Now()

	if result := common.Database.Model(&orm.LogAuctionVehicleStatus{}).Create(log); result.Error != nil {
		infra.AppLog.EventError(result.Error, "SaveLogAuctionVehicleStatus", map[string]any{
			"auction_id":        auctionID,
			"vehicle_id":        vehicleID,
			"auction_status_id": auctionStatusID,
		}, vehicleID,
			infra.WithComponent("auction"),
			infra.WithOperation("save_log_auction_vehicle_status"),
			infra.WithLogKind(infra.LogKindError))
		return result.Error
	}

	return nil
}
