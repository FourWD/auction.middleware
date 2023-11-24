package utils

import (
	"errors"
	"fmt"

	"github.com/FourWD/auction.middleware/model"
	"github.com/FourWD/middleware/common"
	"github.com/gofiber/fiber/v2"
)

func GetJwtToken(c *fiber.Ctx) (*model.AuctionToken, error) {

	err := errors.New("error")
	if err != nil {
		fmt.Println("GetJwtToken")
	}
	auctionToken := new(model.AuctionToken)
	auctionToken.UserID, _ = common.EncodedJwtToken(c, "user_id")
	auctionToken.MobileNo, _ = common.EncodedJwtToken(c, "mobile_no")
	auctionToken.DeviceID, err = common.EncodedJwtToken(c, "noti_token")

	if err == nil {
		return auctionToken, nil
	} else {
		return nil, err
	}
}
