package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type LogUserVip struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	UserID    string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	StartDate time.Time `json:"start_date" query:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" query:"end_date" gorm:"not null"`
	Remark    string    `json:"remark" query:"remark" gorm:"type:text"`
	// IsCancel      bool      `json:"is_cancel" query:"is_cancel" gorm:"bool"`
	// ExpireDate   time.Time `json:"expire_date" query:"expire_date" `
	// CancelBy      string    `json:"cancel_by" query:"cancel_by" gorm:"type:varchar(36)"`
	// CancelRemark  string    `json:"cancel_remark" query:"cancel_remark" gorm:"type:text"`
}
