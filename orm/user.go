package orm

import "github.com/FourWD/middleware/orm"

type User struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserTypeID string `json:"user_type_id" query:"user_type_id" db:"user_type_id" gorm:"type:varchar(36);"`

	Username          string `json:"username" query:"username" db:"username" gorm:"type:varchar(20);"`
	Password          string `json:"password" query:"password" db:"password" gorm:"type:varchar(20);"`
	Mobile            string `json:"mobile" query:"mobile" db:"mobile" gorm:"type:varchar(20);"`
	Email             string `json:"email" query:"email" db:"email" gorm:"type:varchar(20);"`
	Pin               string `json:"pin" query:"pin" db:"pin" gorm:"type:int(6);"`
	NotificationToken string `json:"notification_token" query:"notification_token" db:"notification_token" gorm:"type:varchar(20);"`
}
