package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Auction struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FinanceID string `json:"finance_id" query:"finance_id" gorm:"type:varchar(36)"`

	Code             string    `json:"code" query:"code" gorm:"type:varchar(20)"`
	Name             string    `json:"name" query:"name" gorm:"type:varchar(50)"`
	ShowDate         time.Time `json:"show_date" query:"show_date"`
	StartDate        time.Time `json:"start_date" query:"start_date"`
	EndDate          time.Time `json:"end_date" query:"end_date"`
	CountVehicle     int       `json:"count_vehicle" query:"count_vehicle" gorm:"type:int(6);"`
	CountUserBidding int       `json:"count_user_bidding" query:"count_user_bidding" gorm:"type:int(6);"`
	BiddingStep1     int       `json:"bidding_step_1" query:"bidding_step_1" gorm:"type:int(6);"`
	BiddingStep2     int       `json:"bidding_step_2" query:"bidding_step_2" gorm:"type:int(6);"`
	BiddingStep3     int       `json:"bidding_step_3" query:"bidding_step_3" gorm:"type:int(6);"`
	IsShow           bool      `json:"is_show" query:"is_show" gorm:"bool"`
	IsExtra          bool      `json:"is_extra" query:"is_extra" gorm:"bool"`
	ExtraTimeMinute  int       `json:"extra_time_minute" query:"extra_time_minute" gorm:"type:int(3);"` // 15 min in int
	IsEnd            bool      `json:"is_end" query:"is_end" gorm:"bool"`
}
