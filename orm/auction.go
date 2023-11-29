package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Auction struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); "`
	model.GormModel

	RoundID           string    `json:"round_id" query:"round_id" gorm:"type:varchar(36)"`
	FinanceID         string    `json:"finance_id" query:"finance_id" gorm:"type:varchar(10)"`
	Code              string    `json:"code" query:"code" gorm:"type:varchar(20)"`
	Name              string    `json:"name" query:"name" gorm:"type:varchar(200)"`
	ShowDate          time.Time `json:"show_date" query:"show_date"`
	StartDate         time.Time `json:"start_date" query:"start_date"`
	EndDate           time.Time `json:"end_date" query:"end_date"`
	CountVehicle      int       `json:"count_vehicle" query:"count_vehicle" gorm:"type:int(6)"`
	CountUserBidding  int       `json:"count_user_bidding" query:"count_user_bidding" gorm:"type:int(6)"`
	BiddingStep1      int       `json:"bidding_step_1" query:"bidding_step_1" gorm:"column:bidding_step_1;type:int(6)"`
	BiddingStep2      int       `json:"bidding_step_2" query:"bidding_step_2" gorm:"column:bidding_step_2;type:int(6)"`
	BiddingStep3      int       `json:"bidding_step_3" query:"bidding_step_3" gorm:"column:bidding_step_3;type:int(6)"`
	IsShow            bool      `json:"is_show" query:"is_show" gorm:"bool"`
	IsExtra           bool      `json:"is_extra" query:"is_extra" gorm:"bool"`
	IsProcess         bool      `json:"is_process" query:"is_process" gorm:"bool"`
	ExtraTimeMinute   int       `json:"extra_time_minute" query:"extra_time_minute" gorm:"type:int(3)"` // 15 min in int
	IsEnd             bool      `json:"is_end" query:"is_end" gorm:"bool"`
	EndBy             string    `json:"end_by" query:"end_by" gorm:"type:varchar(36)"`
	IsAutoEnd         bool      `json:"is_auto_end" query:"is_auto_end" gorm:"bool"`
	IsImportRedbook   bool      `json:"is_import_redbook" query:"is_import_redbook" gorm:"bool"`
	ImportRedbookDate time.Time `json:"import_redbook_date" query:"import_redbook_date"`
	ImportRedbookBy   string    `json:"import_redbook_by" query:"import_redbook_by"`
	Description       string    `json:"description" query:"description" gorm:"type:varchar(350)"`
	RunningNo         int       `json:"running_no" query:"running_no" gorm:"primary_key;"`
}
