package utils

import (
	"errors"
	"fmt"

	"github.com/FourWD/auction.middleware/orm"
	"github.com/FourWD/middleware/common"
)

// type AuctionVehicles struct {
// 	ID              string `json:"id" query:"id" firestore:"id"`
// 	RightDeposit    int    `json:"right_deposit" query:"right_deposit" firestore:"right_deposit"`
// 	PaymentTypeID   string `json:"payment_type_id" query:"payment_type_id" firestore:"payment_type_id" `
// 	UserDisplayName string `json:"user_display_name" query:"user_display_name" firestore:"user_display_name"`
// }

func SyncAuctionVehicle(auctionID string, vehicleID string) error {

	if auctionID == "" {
		return errors.New("wrong auction")
	}
	if vehicleID == "" {
		return errors.New("wrong vehicle")
	}

	avehicles := getVehicleFB(auctionID, vehicleID)

	common.Database.Model(orm.AuctionVehicle{}).Debug().Updates(&avehicles)

	return nil
}

func getVehicleFB(auctionID string, vehicleID string) map[string]interface{} {

	dsnap, err := common.FirebaseClient.Collection("auctions").Doc(auctionID).Collection("vehicles").Doc(vehicleID).Get(common.FirebaseCtx)
	if err != nil {
		return nil
	}
	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)
	return m

}
