package orm

import "time"

type TempBidding struct {
	ID uint `json:"id" query:"id" gorm:"primaryKey;autoIncrement"`

	AuctionID string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_temp_bidding"`
	VehicleID string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_temp_bidding"`
	ThreadID  int       `json:"thread_id" query:"thread_id" gorm:"type:int(4);index"`
	IsDone    bool      `json:"is_done" query:"is_done" gorm:"type:bool"`
	DoneDate  time.Time `json:"done_date" query:"done_date"`
}
