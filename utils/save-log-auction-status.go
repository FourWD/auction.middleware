package utils

import (
	"time"

	"github.com/FourWD/auction.middleware/orm"

	"github.com/FourWD/middleware/common"
	"github.com/FourWD/middleware/infra"
	"github.com/google/uuid"
)

func SaveLogAuctionStatus(auctionID string, auctionStatusID string) error {
	log := &orm.LogAuctionStatus{
		ID:              uuid.NewString(),
		AuctionID:       auctionID,
		AuctionStatusID: auctionStatusID,
	}
	log.CreatedAt = time.Now()

	if result := common.Database.Model(&orm.LogAuctionStatus{}).Create(log); result.Error != nil {
		infra.AppLog.EventError(result.Error, "SaveLogAuctionStatus", map[string]any{
			"auction_id":        auctionID,
			"auction_status_id": auctionStatusID,
		}, auctionID,
			infra.WithComponent("auction"),
			infra.WithOperation("save_log_auction_status"),
			infra.WithLogKind(infra.LogKindError))
		return result.Error
	}

	return nil
}
