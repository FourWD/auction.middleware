package orm

import "github.com/FourWD/middleware/orm"

type User struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserTypeID string `json:"user_type_id" query:"user_type_id" db:"user_type_id" gorm:"type:varchar(36);"`

	Username            string `json:"username" query:"username" db:"username" gorm:"type:varchar(20);"`
	Password            string `json:"password" query:"password" db:"password" gorm:"type:varchar(20);"`
	Firstname           string `json:"firstname" query:"firstname" db:"firstname" gorm:"type:varchar(100);"`
	Lastname            string `json:"lastname" query:"lastname" db:"lastname" gorm:"type:varchar(100);"`
	FileAvatarID        string `json:"file_avartar_id" query:"file_avartar_id" db:"file_avartar_id" gorm:"type:varchar(36)"`
	Mobile              string `json:"mobile" query:"mobile" db:"mobile" gorm:"type:varchar(20);"`
	Email               string `json:"email" query:"email" db:"email" gorm:"type:varchar(20);"`
	Pin                 string `json:"pin" query:"pin" db:"pin" gorm:"type:int(6);"`
	IsNotiBeforeAuction bool   `json:"is_noti_before_auction" query:"is_noti_before_auction" db:"is_noti_before_auction" gorm:"type:boolean"`
	IsNotiDuringAuction bool   `json:"is_noti_during_auction" query:"is_noti_during_auction" db:"is_noti_during_auction" gorm:"type:boolean"`
	IsNotiAfterAuction  bool   `json:"is_noti_after_auction" query:"is_noti_after_auction" db:"is_noti_after_auction" gorm:"type:boolean"`
	IsUseTouchID        string `json:"is_use_touch_id" query:"is_use_touch_id" db:"is_use_touch_id" gorm:"type:int(6);"`
	NotificationToken   string `json:"notification_token" query:"notification_token" db:"notification_token" gorm:"type:varchar(20);"`
}
