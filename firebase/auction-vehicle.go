package firebase

import "time"

type AuctionVehicle struct {
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
	ActualEndDate     time.Time `firestore:"actual_end_date"`
	IsEnd             bool      `firestore:"is_end"`
	IsExtra           bool      `firestore:"is_extra"`
	MinPrice          int       `firestore:"min_price"`
	OpenPrice         int       `firestore:"open_price"`
	UserID            string    `firestore:"user_id"`
	BiddingStep1      int       `firestore:"bidding_step_1"`
	BiddingStep2      int       `firestore:"bidding_step_2"`
	BiddingStep3      int       `firestore:"bidding_step_3"`
}
