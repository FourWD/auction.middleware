package utils

import (
	"log"

	"github.com/FourWD/auction.middleware/model"
	"github.com/FourWD/middleware/common"
)

func PublicNotiAuctionEnd(auctionID string) error { // for Aum
	noti := new(model.NotiAuction)
	noti.AuctionID = auctionID

	message := "END@" + common.StructToString(noti)
	if _, err := common.GooglePublicMessage("NOTI", message); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
