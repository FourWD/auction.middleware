package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type User struct {
	orm.User
	RMEmployeeID        string    `json:"rm_employee_id" query:"rm_employee_id" gorm:"type:varchar(36)"`
	Pin                 string    `json:"pin" query:"pin" gorm:"type:varchar(255)"`
	IsNotiBeforeAuction bool      `json:"is_noti_before_auction" query:"is_noti_before_auction" gorm:"type:bool"`
	IsNotiDuringAuction bool      `json:"is_noti_during_auction" query:"is_noti_during_auction" gorm:"type:bool"`
	IsNotiAfterAuction  bool      `json:"is_noti_after_auction" query:"is_noti_after_auction" gorm:"type:bool"`
	IsUseTouchID        bool      `json:"is_use_touch_id" query:"is_use_touch_id" gorm:"type:bool"`
	RightDeposit        int       `json:"right_deposit" query:"right_deposit" gorm:"type:int(5)"`
	ReferralCode        string    `json:"referral_code" query:"referral_code" gorm:"type:varchar(20)"` // รหัสผู้แนะนำ
	IsApprove           bool      `json:"is_approve" query:"is_approve" gorm:"type:bool"`
	ApproveBy           string    `json:"approve_by" query:"approve_by" gorm:"type:varchar(36)"`
	ApproveDate         time.Time `json:"approve_date" query:"approve_dates"`
	PaymentTypeID       string    `json:"payment_type_id" query:"payment_type_id" gorm:"type:varchar(2)"`
}

//	//AuctionCode              string    `json:"auction_code" query:"auction_code" gorm:"type:varchar(20)"` //รหัสผู้ประมูล
//VerifyCode               string    `json:"verify_code" query:"verify_code" gorm:"type:varchar(20)"`   //รหัสยื่นเรื่อง
