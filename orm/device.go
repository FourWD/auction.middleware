package orm

import "github.com/FourWD/middleware/orm"

type Device struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID            string `json:"user_id" query:"user_id"  gorm:"type:varchar(36)"`
	DeviceID          string `json:"device_id" query:"device_id"  gorm:"type:varchar(36)"`
	DeviceName        string `json:"device_name" query:"device_name"  gorm:"type:varchar(255)"`
	Latitude          string `json:"latitude" query:"latitude"  gorm:"type:text"`
	Longitude         string `json:"longitude" query:"longitude"  gorm:"type:text"`
	LocationName      string `json:"location_name" query:"button_url"  gorm:"type:varchar(255)"`
	Token             string `json:"token" query:"token" gorm:"type:varchar(500)"`
	NotificationToken string `json:"notification_token" query:"notification_token" gorm:"type:varchar(20)"`
	IsActive          bool   `json:"is_active" query:"is_active" gorm:"type:bool"`

	orm.RowOrder
}
