package orm

import "github.com/FourWD/middleware/orm"

type Notification struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	ToUserID       string `json:"to_user_id" query:"to_user_id" db:"to_user_id" gorm:"type:varchar(36)"`
	NotificationID string `json:"notification_type_id" query:"notification_type_id" db:"notification_type_id" gorm:"type:varchar(36)"`
	Message        string `json:"message" query:"message" db:"message" gorm:"type:varchar(500)"`
}
