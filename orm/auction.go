package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Auction struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); uniqueIndex:idx_id"`
	model.GormModel

	RoundID            string    `json:"round_id" query:"round_id" gorm:"type:varchar(36)"`
	FinanceID          string    `json:"finance_id" query:"finance_id" gorm:"type:varchar(10)"`
	Code               string    `json:"code" query:"code" gorm:"type:varchar(20)"`
	Name               string    `json:"name" query:"name" gorm:"type:varchar(200)"`
	ShowDate           time.Time `json:"show_date" query:"show_date"`
	StartDate          time.Time `json:"start_date" query:"start_date"`
	EndDate            time.Time `json:"end_date" query:"end_date"`
	CountVehicle       int       `json:"count_vehicle" query:"count_vehicle" gorm:"type:int(6)"`
	CountUserBidding   int       `json:"count_user_bidding" query:"count_user_bidding" gorm:"type:int(6)"`
	BiddingStep1       int       `json:"bidding_step_1" query:"bidding_step_1" gorm:"column:bidding_step_1;type:int(6)"`
	BiddingStep2       int       `json:"bidding_step_2" query:"bidding_step_2" gorm:"column:bidding_step_2;type:int(6)"`
	BiddingStep3       int       `json:"bidding_step_3" query:"bidding_step_3" gorm:"column:bidding_step_3;type:int(6)"`
	AuctionStatusID    string    `json:"auction_status_id" query:"auction_status_id" gorm:"type:varchar(2)"`
	IsExtra            bool      `json:"is_extra" query:"is_extra" gorm:"bool"`
	IsProcess          bool      `json:"is_process" query:"is_process" gorm:"bool"`
	IsShow             bool      `json:"is_show" query:"is_show" gorm:"bool"`
	ExtraTimeMinute    int       `json:"extra_time_minute" query:"extra_time_minute" gorm:"type:int(3)"` // 15 min in int
	IsEnd              bool      `json:"is_end" query:"is_end" gorm:"bool"`
	EndBy              string    `json:"end_by" query:"end_by" gorm:"type:varchar(36)"` //asdsd
	IsAutoEnd          bool      `json:"is_auto_end" query:"is_auto_end" gorm:"bool"`
	IsOperationApprove bool      `json:"is_operation_approve" query:"is_operation_approve" gorm:"bool"`
	OperationBy        string    `json:"operation_by" query:"operation_by" gorm:"type:varchar(36)"`
	IsOwnerApprove     bool      `json:"is_owner_approve" query:"is_owner_approve" gorm:"bool"`
	OwnerBy            string    `json:"owner_by" query:"owner_by" gorm:"type:varchar(36)"`
	IsImportRedbook    bool      `json:"is_import_redbook" query:"is_import_redbook" gorm:"bool"`
	ImportRedbookDate  time.Time `json:"import_redbook_date" query:"import_redbook_date"`
	ImportRedbookBy    string    `json:"import_redbook_by" query:"import_redbook_by"`
	Description        string    `json:"description" query:"description" gorm:"type:varchar(350)"`

	IsAllVehicleApprove bool `json:"is_all_vehicle_approve" query:"is_all_vehicle_approve" gorm:"bool"`

	RunningNo int `json:"running_no" query:"running_no" gorm:"primary_key;auto_increment;not_null"`

	CountOfSedan  int `json:"count_of_sedan" query:"count_of_sedan" gorm:"type:int(6)"`
	CountOfPickUp int `json:"count_of_pick_up" query:"count_of_pick_up" gorm:"type:int(6)"`
	CountOfSUV    int `json:"count_of_suv" query:"count_of_suv" gorm:"type:int(6)"`
	CountOfVan    int `json:"count_of_van" query:"count_of_van" gorm:"type:int(6)"`
	CountOfTruck  int `json:"count_of_truck" query:"count_of_truck" gorm:"type:int(6)"`
}
