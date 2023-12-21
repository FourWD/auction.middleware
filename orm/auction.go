package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Auction struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); uniqueIndex:idx_id"`
	model.GormModel

	RoundID            string    `json:"round_id" query:"round_id" gorm:"type:varchar(36)"`
	Code               string    `json:"code" query:"code" gorm:"type:varchar(20)"`
	Name               string    `json:"name" query:"name" gorm:"type:varchar(200)"`
	ShowDate           time.Time `json:"show_date" query:"show_date"`
	StartDate          time.Time `json:"start_date" query:"start_date"`
	EndDate            time.Time `json:"end_date" query:"end_date"`
	ActualEndDate      time.Time `json:"actual_end_date" query:"actual_end_date"`
	CountVehicle       int       `json:"count_vehicle" query:"count_vehicle" gorm:"type:int(6)"`
	//CountBidding       int       `json:"count_bidding" query:"count_bidding" gorm:"type:int(6)"`
	//CountProxy         int       `json:"count_proxy" query:"count_proxy" gorm:"type:int(6)"`
	CountUserBidding   int       `json:"count_user_bidding" query:"count_user_bidding" gorm:"type:int(6)"`
	//CountUserProxy     int       `json:"count_user_proxy" query:"count_user_proxy" gorm:"type:int(6)"`
	//CountUserFavorite  int       `json:"count_user_favorite" query:"count_user_favorite" gorm:"type:int(6)"`
	//CountUserJoin      int       `json:"count_user_join" query:"count_user_join" gorm:"type:int(6)"`
	//CountUserView      int       `json:"count_user_view" query:"count_user_view" gorm:"type:int(6)"`
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
	CountOfPickup int `json:"count_of_pickup" query:"count_of_pickup" gorm:"type:int(6)"` // char U to u
	CountOfSUV    int `json:"count_of_suv" query:"count_of_suv" gorm:"type:int(6)"`
	CountOfVan    int `json:"count_of_van" query:"count_of_van" gorm:"type:int(6)"`
	CountOfTruck  int `json:"count_of_truck" query:"count_of_truck" gorm:"type:int(6)"`
}
