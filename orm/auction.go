package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Auction struct {
	ID string `json:"id" query:"id" firestore:"id" gorm:"type:varchar(36); uniqueIndex:idx_id"`
	model.GormModel

	RoundID              string    `json:"round_id" query:"round_id" firestore:"round_id" gorm:"type:varchar(36)"`
	Code                 string    `json:"code" query:"code" firestore:"code" gorm:"type:varchar(20)"`
	Name                 string    `json:"name" query:"name" firestore:"name" gorm:"type:varchar(200)"`
	ShowDate             time.Time `json:"show_date" query:"show_date" firestore:"show_date"`
	StartDate            time.Time `json:"start_date" query:"start_date" firestore:"start_date"`
	EndDate              time.Time `json:"end_date" query:"end_date" firestore:"end_date"`
	ActualEndDate        time.Time `json:"actual_end_date" query:"actual_end_date" firestore:"actual_end_date"`
	CountVehicle         int       `json:"count_vehicle" query:"count_vehicle" firestore:"count_vehicle" gorm:"type:int(6)"`
	CountUserJoin        int       `json:"count_user_join" query:"count_user_join" firestore:"count_user_join" gorm:"type:int(6)"`
	CountUserBidding     int       `json:"count_user_bidding" query:"count_user_bidding" firestore:"count_user_bidding" gorm:"type:int(6)"`
	AuctionStatusID      string    `json:"auction_status_id" query:"auction_status_id" firestore:"auction_status_id" gorm:"type:varchar(2)"`
	IsExtra              bool      `json:"is_extra" query:"is_extra" firestore:"is_extra" gorm:"type:bool"`
	IsShow               bool      `json:"is_show" query:"is_show" firestore:"is_show" gorm:"type:bool"`
	IsMinPrice           bool      `json:"is_min_price" query:"is_min_price" firestore:"is_min_price" gorm:"type:bool"`
	ExtraTimeMinute      int       `json:"extra_time_minute" query:"extra_time_minute" firestore:"extra_time_minute" gorm:"type:int(3)"`
	IsApprove            bool      `json:"is_approve" query:"is_approve" firestore:"is_approve" gorm:"type:bool"`
	IsAutoApprove        bool      `json:"is_auto_approve" query:"is_auto_approve" firestore:"is_auto_approve" gorm:"type:bool"`
	IsOperationApprove   bool      `json:"is_operation_approve" query:"is_operation_approve" firestore:"is_operation_approve" gorm:"type:bool"`
	OperationApproveBy   string    `json:"operation_approve_by" query:"operation_approve_by" firestore:"operation_approve_by" gorm:"type:varchar(36)"`
	OperationApproveDate time.Time `json:"operation_approve_date" query:"operation_approve_date" firestore:"operation_approve_date"`
	IsOwnerApprove       bool      `json:"is_owner_approve" query:"is_owner_approve" firestore:"is_owner_approve" gorm:"type:bool"`
	OwnerApproveBy       string    `json:"owner_approve_by" query:"owner_approve_by" firestore:"owner_approve_by" gorm:"type:varchar(36)"`
	OwnerApproveDate     time.Time `json:"owner_approve_date" query:"owner_approve_date" firestore:"owner_approve_date"` //วันที่ประกาศผล
	IsImportRedbook      bool      `json:"is_import_redbook" query:"is_import_redbook" firestore:"is_import_redbook" gorm:"type:bool"`
	ImportRedbookDate    time.Time `json:"import_redbook_date" query:"import_redbook_date" firestore:"import_redbook_date" `
	ImportRedbookBy      string    `json:"import_redbook_by" query:"import_redbook_by" firestore:"import_redbook_by" gorm:"type:varchar(36)"`
	Description          string    `json:"description" query:"description" firestore:"description" gorm:"type:varchar(350)"`
	IsAllVehicleApprove  bool      `json:"is_all_vehicle_approve" query:"is_all_vehicle_approve" firestore:"is_all_vehicle_approve" gorm:"type:bool"`
	IsGenFile            bool      `json:"is_gen_file" query:"is_gen_file" firestore:"is_gen_file" gorm:"type:bool"`
	GenFileDate          time.Time `json:"gen_file_date" query:"gen_file_date" firestore:"gen_file_date" `

	RunningNo     int  `json:"running_no" query:"running_no" firestore:"running_no" gorm:"primary_key;auto_increment;not_null"`
	CountOfSedan  int  `json:"count_of_sedan" query:"count_of_sedan" firestore:"count_of_sedan" gorm:"type:int(6)"`
	CountOfPickup int  `json:"count_of_pickup" query:"count_of_pickup" firestore:"count_of_pickup" gorm:"type:int(6)"`
	CountOfSUV    int  `json:"count_of_suv" query:"count_of_suv" firestore:"count_of_suv" gorm:"type:int(6)"`
	CountOfVan    int  `json:"count_of_van" query:"count_of_van" firestore:"count_of_van" gorm:"type:int(6)"`
	CountOfTruck  int  `json:"count_of_truck" query:"count_of_truck" firestore:"count_of_truck" gorm:"type:int(6)"`
	IsSyncBidding bool `json:"is_sync_bidding" query:"is_sync_bidding" firestore:"is_sync_bidding" gorm:"type:bool"`
	BiddingStep1  int  `json:"bidding_step_1" query:"bidding_step_1" firestore:"bidding_step_1" gorm:"column:bidding_step_1;type:int(6)"`
	BiddingStep2  int  `json:"bidding_step_2" query:"bidding_step_2" firestore:"bidding_step_2" gorm:"column:bidding_step_2;type:int(6)"`
	BiddingStep3  int  `json:"bidding_step_3" query:"bidding_step_3" firestore:"bidding_step_3" gorm:"column:bidding_step_3;type:int(6)"`
}

/*
fourwd approve button
update is_operation_approve = 1 ,
 operation_approve_by = jwtuser ,
 operation_approve_date

oma approve button
update is_operation_approve = 1 ,
owner_approve_by = jwtuser ,
owner_approve_date ,
 is_approve = 1 ,
 auction_status_id = "05"
*/
