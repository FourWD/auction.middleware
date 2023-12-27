package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type AuctionVehicle struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AuctionID          string    `json:"auction_id" query:"auction_id" firestore:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`
	VehicleID          string    `json:"vehicle_id" query:"vehicle_id" firestore:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`
	VehicleNo          string    `json:"vehicle_no" query:"vehicle_no" firestore:"vehicle_no" gorm:"type:varchar(10); "`
	OpenPrice          int       `json:"open_price" query:"open_price" firestore:"open_price" gorm:"type:int(10)"`
	MinPrice           int       `json:"min_price" query:"min_price" firestore:"min_price" gorm:"type:int(10)"`
	ClosePrice         int       `json:"close_price" query:"close_price" firestore:"close_price" gorm:"type:int(10)"`
	IsWin              bool      `json:"is_win" query:"is_win" firestore:"is_win" gorm:"type:bool"`
	StartDate          time.Time `json:"start_date" query:"start_date" firestore:"start_date"`
	EndDate            time.Time `json:"end_date" query:"end_date" firestore:"end_date"`
	ActualEndDate      time.Time `json:"actual_end_date" query:"actual_end_date" firestore:"actual_end_date"`
	LastBiddingDate    time.Time `json:"last_bidding_date" query:"last_bidding_date" firestore:"last_bidding_date"`
	LastProcessDate    time.Time `json:"last_process_date" query:"last_process_date" firestore:"last_process_date"`
	IsExtra            bool      `json:"is_extra" query:"is_extra" firestore:"is_extra" gorm:"bool"`
	IsEnd              bool      `json:"is_end" query:"is_end" firestore:"is_end" gorm:"bool"`
	IsOperationApprove bool      `json:"is_operation_approve" query:"is_operation_approve" firestore:"is_operation_approve" gorm:"bool"`
	IsApprove          bool      `json:"is_approve" query:"is_approve" firestore:"is_approve" gorm:"bool"`
	OperationApproveBy string    `json:"operation_approve_by" query:"operation_approve_by" firestore:"operation_approve_by" gorm:"type:varchar(36)"`
	BiddingStep1       int       `json:"bidding_step_1" query:"bidding_step_1" firestore:"bidding_step_1" gorm:"column:bidding_step_1;type:int(6)"`
	BiddingStep2       int       `json:"bidding_step_2" query:"bidding_step_2" firestore:"bidding_step_2" gorm:"column:bidding_step_2;type:int(6)"`
	BiddingStep3       int       `json:"bidding_step_3" query:"bidding_step_3" firestore:"bidding_step_3" gorm:"column:bidding_step_3;type:int(6)"`
	CountUserView      int       `json:"count_user_view" query:"count_user_view" firestore:"count_user_view" gorm:"type:int(10)"`
	CountUserFavorite  int       `json:"count_user_favorite" query:"count_user_favorite" firestore:"count_user_favorite" gorm:"type:int"`
	CountUserJoin      int       `json:"count_user_join" query:"count_user_join" firestore:"count_user_join" gorm:"type:int(10)"`
	CountUserBidding   int       `json:"count_user_bidding" query:"count_user_bidding" firestore:"count_user_bidding"  gorm:"type:int(10)"`
	CountUserProxy     int       `json:"count_user_proxy" query:"count_user_proxy" firestore:"count_user_proxy" gorm:"type:int(10)"`
	CountBidding       int       `json:"count_bidding" query:"count_bidding" firestore:"count_bidding" gorm:"type:int(10)"`
	CountView          int       `json:"count_view" query:"count_view" firestore:"count_view" gorm:"type:int"`
	CountProxy         int       `json:"count_proxy" query:"count_proxy" firestore:"count_proxy" gorm:"type:int"`
	WinnerUserID       string    `json:"winner_user_id" query:"winner_user_id" firestore:"winner_user_id" gorm:"type:varchar(36); "`
	IsWinByProxy       bool      `json:"is_win_by_proxy" query:"is_win_by_proxy" firestore:"is_win_by_proxy" gorm:"bool"`
	CurrentPrice       int       `json:"current_price" query:"current_price" firestore:"current_price" gorm:"type:int"`
	ExtraTimeMinute    int       `json:"firestore" query:"firestore" firestore:"extra_time_minute" gorm:"type:int"`
	RowOrder           float32   `json:"row_order" query:"row_order" firestore:"row_order" gorm:"type:decimal(5,2)"`
}

/* type AuctionVehicle struct {
	CountBidding      int       `firestore:"count_bidding"`
	CountProxy        int       `firestore:"count_proxy"`
	CountUserFavorite int       `firestore:"count_user_favorite"`
	CountUserView     int       `firestore:"count_user_view"`
	CountUserJoin     int       `firestore:"count_user_join"`
	CountUserBidding  int       `firestore:"count_user_bidding"`
	CountUserProxy    int       `firestore:"count_user_proxy"`
	CurrentPrice      int       `firestore:"current_price"`
	StartDate         time.Time `firestore:"start_date"`
	EndDate           time.Time `firestore:"end_date"`
	ExtraTimeMinute   int       `firestore:"extra_time_minute"`
	ActualEndDate     time.Time `firestore:"actual_end_date"`
	IsEnd             bool      `firestore:"is_end"`
	IsExtra           bool      `firestore:"is_extra"`
	MinPrice          int       `firestore:"min_price"`
	OpenPrice         int       `firestore:"open_price"`
	UserID            string    `firestore:"user_id"`
	BiddingStep1      int       `firestore:"bidding_step_1"`
	BiddingStep2      int       `firestore:"bidding_step_2"`
	BiddingStep3      int       `firestore:"bidding_step_3"`
	CountView         int       `firestore:"count_view"`
}
*/
