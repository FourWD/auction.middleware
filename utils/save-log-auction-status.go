package utils

import (
	"time"

	"github.com/FourWD/auction.middleware/orm"

	"github.com/FourWD/middleware/common"
	"github.com/google/uuid"
)

func SaveLogAuctionStatus(auctionID string, auctionStatusID string) error {
	log := new(orm.LogAuctionStatus)
	log.ID = uuid.NewString()
	log.AuctionID = auctionID
	log.AuctionStatusID = auctionStatusID
	log.CreatedAt = time.Now()
	// ========================================================================================
	logFields := map[string]interface{}{
		"auction_id":        auctionID,
		"auction_status_id": auctionStatusID,
		"created_at":        log.CreatedAt,
	}
	// ========================================================================================
	if err := common.Database.Model(&orm.LogAuctionStatus{}).Create(&log); err.Error != nil {
		logFields["error"] = err.Error
		common.LogError("SaveLogAuctionStatus", logFields, auctionID)
		return err.Error
	}
	// ========================================================================================
	common.Log("SaveLogAuctionStatus", logFields, auctionID)
	return nil
}
