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

	if err := common.Database.Model(&orm.LogAuctionStatus{}).Create(&log); err.Error != nil {
		return err.Error
	}

	return nil
}
