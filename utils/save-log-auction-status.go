package utils

import (
	"time"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
	"github.com/google/uuid"
)

func LogAuctionStatus(auctionID string, auctionStatus string) error {

	logAuctionStatus := new(orm.LogAuctionStatus)

	logAuctionStatus.ID = uuid.NewString()
	logAuctionStatus.AuctionID = auctionID
	logAuctionStatus.AuctionStatusID = auctionStatus
	logAuctionStatus.CreatedAt = time.Now()

	if err := common.Database.Model(&orm.LogAuctionStatus{}).Create(&logAuctionStatus); err.Error != nil {
		return err.Error
	}

	return nil

}
