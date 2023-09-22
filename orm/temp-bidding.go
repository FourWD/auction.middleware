package orm

type TempBidding struct {
	ID uint `json:"id" query:"id" gorm:"primaryKey;autoIncrement"`

	FinanceID    string `json:"finance_id" query:"finance_id" gorm:"type:varchar(36);uniqueIndex:idx_temp_bidding"`
	AuctionID    string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_temp_bidding"`
	VehicleID    string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);uniqueIndex:idx_temp_bidding"`
	Price        int    `json:"price" query:"price" gorm:"type:int(4);uniqueIndex:idx_temp_bidding"`
	ThreadID     int    `json:"thread_id" query:"thread_id" gorm:"type:int(4);index"`
	IsDone       bool   `json:"is_done" query:"is_done" gorm:"type:bool"`
	DoneDate     string `json:"done_date" query:"done_date" gorm:"type:varchar(20)"`
	WinnerUserID string `json:"winner_user_id" query:"winner_user_id" gorm:"type:varchar(36)"`
	WinnerPrice  int    `json:"winner_price" query:"winner_price" gorm:"type:int(4);"`
}
