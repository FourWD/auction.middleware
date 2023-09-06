package utils

import (
	"errors"
	"fmt"

	"github.com/FourWD/middleware-auction/model"
	"github.com/FourWD/middleware/common"
	"github.com/gofiber/fiber/v2"
)

func GetJwtToken(c *fiber.Ctx) (*model.AuctionToken, error) {
	bearerToken := c.Get("Authorization")

	err := errors.New("error")
	if err != nil {
		fmt.Println("GetJwtToken")
	}
	auctionToken := new(model.AuctionToken)
	auctionToken.UserID, err = common.EncodedJwtToken(bearerToken, "user_id")

	if err == nil {
		return auctionToken, nil
	} else {
		return nil, err
	}
}
