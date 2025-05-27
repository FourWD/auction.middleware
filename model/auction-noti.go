package model

type AuctionNoti struct {
	AuctionID string `json:"auction_id" query:"auction_id"`
	VehicleID string `json:"vehicle_id" query:"vehicle_id"`
}
