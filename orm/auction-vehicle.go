package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type AuctionVehicle struct {
	ID string `json:"id" query:"id" firestore:"id" bson:"_id,omitempty" gorm:"type:varchar(36);primary_key"`
	model.GormModel
	AuctionID             string    `json:"auction_id" query:"auction_id" firestore:"auction_id" bson:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`
	VehicleID             string    `json:"vehicle_id" query:"vehicle_id" firestore:"vehicle_id" bson:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_vehicle_id"`
	VehicleNo             int       `json:"vehicle_no" query:"vehicle_no" firestore:"vehicle_no" bson:"vehicle_no" gorm:"type:int; "`
	OpenPrice             int       `json:"open_price" query:"open_price" firestore:"open_price" bson:"open_price" gorm:"type:int(10)"`
	MinPrice              int       `json:"min_price" query:"min_price" firestore:"min_price" bson:"min_price" gorm:"type:int(10)"`
	ClosePrice            int       `json:"close_price" query:"close_price" firestore:"close_price" bson:"close_price" gorm:"type:int(10)"`
	IsWin                 bool      `json:"is_win" query:"is_win" firestore:"is_win" bson:"is_win" gorm:"type:bool"`
	StartDate             time.Time `json:"start_date" query:"start_date" firestore:"start_date" bson:"start_date"`
	EndDate               time.Time `json:"end_date" query:"end_date" firestore:"end_date" bson:"end_date"`
	ActualEndDate         time.Time `json:"actual_end_date" query:"actual_end_date" firestore:"actual_end_date" bson:"actual_end_date"`
	LastBiddingDate       time.Time `json:"last_bidding_date" query:"last_bidding_date" firestore:"last_bidding_date" bson:"last_bidding_date"`
	LastProcessDate       time.Time `json:"last_process_date" query:"last_process_date" firestore:"last_process_date" bson:"last_process_date"`
	IsSyncBidding         bool      `json:"is_sync_bidding" query:"is_sync_bidding" firestore:"is_sync_bidding" bson:"is_sync_bidding" gorm:"type:bool"`
	SyncBiddingDate       time.Time `json:"sync_bidding_date" query:"sync_bidding_date" firestore:"sync_bidding_date" bson:"sync_bidding_date"`
	IsExtra               bool      `json:"is_extra" query:"is_extra" firestore:"is_extra" bson:"is_extra" gorm:"type:bool"`
	IsEnd                 bool      `json:"is_end" query:"is_end" firestore:"is_end" bson:"is_end" gorm:"type:bool"`
	IsApprove             bool      `json:"is_approve" query:"is_approve" firestore:"is_approve" bson:"is_approve" gorm:"type:bool"`
	IsOperationApprove    bool      `json:"is_operation_approve" query:"is_operation_approve" firestore:"is_operation_approve" bson:"is_operation_approve" gorm:"type:bool"`
	OperationApproveBy    string    `json:"operation_approve_by" query:"operation_approve_by" firestore:"operation_approve_by" bson:"operation_approve_by" gorm:"type:varchar(36)"`
	OperationApproveDate  time.Time `json:"operation_approve_date" query:"operation_approve_date" firestore:"operation_approve_date" bson:"operation_approve_date"`
	BiddingStep1          int       `json:"bidding_step_1" query:"bidding_step_1" firestore:"bidding_step_1" bson:"bidding_step_1" gorm:"column:bidding_step_1;type:int(6)"`
	BiddingStep2          int       `json:"bidding_step_2" query:"bidding_step_2" firestore:"bidding_step_2" bson:"bidding_step_2" gorm:"column:bidding_step_2;type:int(6)"`
	BiddingStep3          int       `json:"bidding_step_3" query:"bidding_step_3" firestore:"bidding_step_3" bson:"bidding_step_3" gorm:"column:bidding_step_3;type:int(6)"`
	CountUserView         int       `json:"count_user_view" query:"count_user_view" firestore:"count_user_view" bson:"count_user_view" gorm:"type:int"`
	CountUserFavorite     int       `json:"count_user_favorite" query:"count_user_favorite" firestore:"count_user_favorite" bson:"count_user_favorite" gorm:"type:int"`
	CountUserJoin         int       `json:"count_user_join" query:"count_user_join" firestore:"count_user_join" bson:"count_user_join" gorm:"type:int"`
	CountUserBidding      int       `json:"count_user_bidding" query:"count_user_bidding" firestore:"count_user_bidding" bson:"count_user_bidding" gorm:"type:int"`
	CountUserProxy        int       `json:"count_user_proxy" query:"count_user_proxy" firestore:"count_user_proxy" bson:"count_user_proxy" gorm:"type:int"`
	CountBidding          int       `json:"count_bidding" query:"count_bidding" firestore:"count_bidding" bson:"count_bidding" gorm:"type:int"`
	CountView             int       `json:"count_view" query:"count_view" firestore:"count_view" bson:"count_view" gorm:"type:int"`
	CountProxy            int       `json:"count_proxy" query:"count_proxy" firestore:"count_proxy" bson:"count_proxy" gorm:"type:int"`
	WinnerUserID          string    `json:"winner_user_id" query:"winner_user_id" firestore:"winner_user_id" bson:"winner_user_id" gorm:"type:varchar(36); "`
	WinnerUserAuctionCode string    `json:"winner_user_auction_code" query:"winner_user_auction_code" firestore:"winner_user_auction_code" bson:"winner_user_auction_code" gorm:"type:varchar(5); "`
	WinnerUserDisplayName string    `json:"winner_user_display_name" query:"winner_user_display_name" firestore:"winner_user_display_name" bson:"winner_user_display_name" gorm:"type:varchar(500)"`
	IsWinByProxy          bool      `json:"is_win_by_proxy" query:"is_win_by_proxy" firestore:"is_win_by_proxy" bson:"is_win_by_proxy" gorm:"type:bool"`
	IsProcess             bool      `json:"is_process" query:"is_process" firestore:"is_process" bson:"is_process" gorm:"type:bool"`
	CurrentPrice          int       `json:"current_price" query:"current_price" firestore:"current_price" bson:"current_price" gorm:"type:int"`
	CurrentProxyID        string    `json:"current_proxy_id" query:"current_proxy_id" firestore:"current_proxy_id" bson:"current_proxy_id" gorm:"type:varchar(36)"`
	ExtraTimeMinute       int       `json:"extra_time_minute" query:"extra_time_minute" firestore:"extra_time_minute" bson:"extra_time_minute" gorm:"type:int"`
	AuctionStatusID       string    `json:"auction_status_id" query:"auction_status_id" firestore:"auction_status_id" bson:"auction_status_id" gorm:"type:varchar(2)"`
	FirstOpenPrice        int       `json:"first_open_price" query:"first_open_price" firestore:"first_open_price" bson:"first_open_price" gorm:"type:int"`
	CountAuction          int       `json:"count_auction" query:"count_auction" firestore:"count_auction" bson:"count_auction" gorm:"type:int"` // ผ่านการประมูลมาแล้วกี่รอบ
	ProxyToken            string    `json:"proxy_token" query:"proxy_token" firestore:"proxy_token" bson:"proxy_token" gorm:"type:varchar(36);default:NULL"`
	RowOrder              float32   `json:"row_order" query:"row_order" firestore:"row_order" bson:"row_order" gorm:"type:decimal(5,2)"`
	// LogLastBiddingPrice   int       `json:"log_last_bidding_price" query:"log_last_bidding_price" firestore:"log_last_bidding_price" gorm:"type:int"`
	// LogLastBiddingUserID  string    `json:"log_last_bidding_user_id" query:"log_last_bidding_user_id" firestore:"log_last_bidding_user_id" gorm:"type:varchar(36)"`
	// LogMaxProxyID         string    `json:"log_max_proxy_id" query:"log_max_proxy_id" firestore:"log_max_proxy_id" gorm:"type:varchar(36)"`
	// LogMaxProxyPrice      int       `json:"log_max_proxy_price" query:"log_max_proxy_price" firestore:"log_max_proxy_price" gorm:"type:int"`
	// LogMaxProxyUserID     string    `json:"log_max_proxy_user_id" query:"log_max_proxy_user_id" firestore:"log_max_proxy_user_id" gorm:"type:varchar(36)"` // comment model p'del for create auction vehicle
}

/* //


operation approve button
update is_operation_approve = 1 ,
operation_approve_by = jwtuser ,
operation_approve_date ,
 is_approve = 1 ,
 auction_status_id = "05"

 ปุ่มนี้กดได้ is_win = true ด้วย
*/
