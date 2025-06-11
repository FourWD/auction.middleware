package utils

import (
	"time"

	"github.com/FourWD/auction.middleware/model"
)

func IsNotiAuctionExpire(noti model.NotiAuction) bool {
	expire := noti.ActionTime.Add(5 * time.Minute)
	return time.Now().After(expire)
}
