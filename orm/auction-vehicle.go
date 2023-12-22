package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type AuctionVehicle struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AuctionID          string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`
	VehicleID          string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`
	VehicleNo          string    `json:"vehicle_no" query:"vehicle_no" gorm:"type:varchar(10); "`
	OpenPrice          int       `json:"open_price" query:"open_price" gorm:"type:int(10)"`
	MinPrice           int       `json:"min_price" query:"min_price" gorm:"type:int(10)"`
	ClosePrice         int       `json:"close_price" query:"close_price" gorm:"type:int(10)"`
	IsWin              bool      `json:"is_win" query:"is_win" gorm:"type:bool"`
	StartDate          time.Time `json:"start_date" query:"start_date"`
	EndDate            time.Time `json:"end_date" query:"end_date"`
	ActualEndDate      time.Time `json:"actual_end_date" query:"actual_end_date"`
	IsExtra            bool      `json:"is_extra" query:"is_extra" gorm:"bool"`
	IsOperationApprove bool      `json:"is_operation_approve" query:"is_operation_approve" gorm:"bool"`
	IsApprove          bool      `json:"is_approve" query:"is_approve" gorm:"bool"`
	OperationBy        string    `json:"operation_by" query:"operation_by" gorm:"type:varchar(36)"`
	BiddingStep1       int       `json:"bidding_step_1" query:"bidding_step_1" gorm:"column:bidding_step_1;type:int(6)"`
	BiddingStep2       int       `json:"bidding_step_2" query:"bidding_step_2" gorm:"column:bidding_step_2;type:int(6)"`
	BiddingStep3       int       `json:"bidding_step_3" query:"bidding_step_3" gorm:"column:bidding_step_3;type:int(6)"`
	CountUserView      int       `json:"count_user_view" query:"count_user_view" gorm:"type:int(10)"`
	CountUserFavorite  int       `json:"count_user_favorite" query:"count_user_favorite" gorm:"type:int"`
	CountUserJoin      int       `json:"count_user_join" query:"count_user_join" gorm:"type:int(10)"`
	CountUserBidding   int       `json:"count_user_bidding" query:"count_user_bidding" gorm:"type:int(10)"`
	CountUserProxy     int       `json:"count_user_proxy" query:"count_user_proxy" gorm:"type:int(10)"`
	CountBidding       int       `json:"count_bidding" query:"count_bidding" gorm:"type:int(10)"`
	CountView          int       `json:"count_view" query:"count_view" gorm:"type:int"`
	WinnerUserID       string    `json:"winner_user_id" query:"winner_user_id" gorm:"type:varchar(36); "`
	IsWinByProxy       bool      `json:"is_win_by_proxy" query:"vehicle_no" gorm:"type:varchar(10); "`
	RowOrder           float32   `json:"row_order" query:"row_order" gorm:"type:decimal(5,2)"`
}
