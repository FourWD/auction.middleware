package model

type NotiAuction struct {
	AuctionID string `json:"auction_id" query:"auction_id"`
	VehicleID string `json:"vehicle_id" query:"vehicle_id"`
	UserID    string `json:"user_id" query:"user_id"`
}
