package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

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

	avehicles, _ := getVehicleFB(auctionID, vehicleID)

	updateData := map[string]interface{}{
		"count_user_view":          avehicles.CountUserView,
		"count_user_join":          avehicles.CountUserJoin,
		"count_user_bidding":       avehicles.CountUserBidding,
		"count_user_proxy":         avehicles.CountUserProxy,
		"count_user_favorite":      avehicles.CountUserFavorite,
		"count_view":               avehicles.CountView,
		"count_bidding":            avehicles.CountBidding,
		"count_proxy":              avehicles.CountProxy,
		"is_win":                   avehicles.IsWin,
		"close_price":              avehicles.ClosePrice,
		"current_price":            avehicles.CurrentPrice,
		"winner_user_id":           avehicles.WinnerUserID,
		"winner_user_auction_code": avehicles.WinnerUserAuctionCode,
		"winner_user_display_name": avehicles.WinnerUserDisplayName,
		"actual_end_date":          avehicles.ActualEndDate,
		"is_win_by_proxy":          avehicles.IsWinByProxy,
		"is_extra":                 avehicles.IsExtra,
	}
	// if err := common.Database.Model(&orm.Auction{}).Where("id = ?", auctionID).Updates(updateData).Error; err != nil {
	// 	return err
	// }
	if err := common.Database.Model(&orm.AuctionVehicle{}).Where("auction_id = ? AND vehicle_id = ? ", auctionID, vehicleID).Updates(&updateData).Error; err != nil {
		return err
	}

	return nil
}

func getVehicleFB(auctionID string, vehicleID string) (orm.AuctionVehicle, error) {

	logDesc := ""

	docPath := fmt.Sprintf("auctions/%s/vehicles/%s", auctionID, vehicleID)
	docRef := common.FirebaseClient.Doc(docPath)
	snap, err := docRef.Get(context.Background())
	if err != nil {
		logDesc += "\n Log: cannot get vehicle"
		common.PrintError("GetAuctionVehicleFirebase", logDesc)
		return orm.AuctionVehicle{}, errors.New("cannot get vehicle")
	}

	var vehicle orm.AuctionVehicle
	if err := snap.DataTo(&vehicle); err != nil {
		logDesc += "\n Log: cannot get vehicle"
		common.PrintError("GetAuctionVehicleFirebase", logDesc)
		return orm.AuctionVehicle{}, errors.New("cannot get vehicle")
	}
	// startDate := snap.Data()["start_date"].(time.Time)
	// endDate := snap.Data()["end_date"].(time.Time)
	actualEndDate := snap.Data()["actual_end_date"].(time.Time)
	vehicle.ActualEndDate = common.UTCToThailandTime(actualEndDate.Truncate(time.Second))

	// dsnap, err := common.FirebaseClient.Collection("auctions").Doc(auctionID).Collection("vehicles").Doc(vehicleID).Get(common.FirebaseCtx)
	// if err != nil {
	// 	return nil
	// }
	// m := dsnap.Data()
	// fmt.Printf("Document data: %#v\n", m)
	return vehicle, nil

}
