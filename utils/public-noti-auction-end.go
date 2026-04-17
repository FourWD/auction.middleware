package utils

import (
	"time"

	"github.com/FourWD/auction.middleware/model"
	"github.com/FourWD/middleware/infra"
	"github.com/FourWD/middleware/kit"
)

func PublicNotiAuctionEnd(auctionID string, actionTime time.Time) error { // for Aum
	noti := new(model.NotiAuction)
	noti.AuctionID = auctionID
	noti.ActionTime = actionTime

	message := "END@" + kit.StructToString(noti)
	if _, err := infra.GooglePublicMessage("NOTI", message); err != nil {
		// log.Println(err)
		return err
	}
	return nil
}
