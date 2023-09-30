package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type User struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Code         string `json:"code" query:"code" gorm:"type:varchar(10)"`
	UserTypeID   string `json:"user_type_id" query:"user_type_id" gorm:"type:varchar(2);"`
	RmEmployeeID string `json:"rm_employee_id" query:"rm_employee_id" gorm:"type:varchar(36);"`

	Username            string    `json:"username" query:"username" gorm:"type:varchar(20);"`
	Password            string    `json:"password" query:"password" gorm:"type:varchar(20);"`
	PrefixID            int       `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2);"`
	Firstname           string    `json:"firstname" query:"firstname" gorm:"type:varchar(100);"`
	Lastname            string    `json:"lastname" query:"lastname" gorm:"type:varchar(100);"`
	FileAvatarID        string    `json:"file_avartar_id" query:"file_avartar_id" gorm:"type:varchar(36)"`
	Mobile              string    `json:"mobile" query:"mobile" gorm:"type:varchar(20); unique"`
	Email               string    `json:"email" query:"email" gorm:"type:varchar(20);"`
	Pin                 string    `json:"pin" query:"pin" gorm:"type:varchar(255);"`
	IsNotiBeforeAuction bool      `json:"is_noti_before_auction" query:"is_noti_before_auction" gorm:"type:bool"`
	IsNotiDuringAuction bool      `json:"is_noti_during_auction" query:"is_noti_during_auction" gorm:"type:bool"`
	IsNotiAfterAuction  bool      `json:"is_noti_after_auction" query:"is_noti_after_auction" gorm:"type:bool"`
	IsUseTouchID        string    `json:"is_use_touch_id" query:"is_use_touch_id" gorm:"type:int(6);"`
	CountWrongLogin     int       `json:"count_wrong_login" query:"count_wrong_login" gorm:"type:int(1);"`
	LastLoginDate       time.Time `json:"last_login_date" query:"last_login_date"`
	UserStatus          string    `json:"user_status" query:"user_status" gorm:"type:varchar(15)"`                   // ban approv
	UserRegisterStatus  string    `json:"user_register_status" query:"user_register_status" gorm:"type:varchar(15)"` //สถานะหน้าสมัคร ถึงขันไหนละ 1 otp เสร็จ 2 กรอกข้อมูลเสร็จ 3 แอดมินแอพพรูฟ
	RightDeposit        int       `json:"right_deposit" query:"right_deposit" gorm:"type:int(5)"`
	ReferralCode        string    `json:"referral_code" query:"referral_code" gorm:"type:varchar(20)"` // รหัสผู้แนะนำ
	IsApprove           bool      `json:"is_approve" query:"is_approve" gorm:"type:bool"`
	ApproveBy           string    `json:"approve_by" query:"approve_by" gorm:"type:varchar(36);"`
	ApproveAt           time.Time `json:"approve_at"`
}
