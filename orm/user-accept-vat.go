package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type UserAcceptVat struct {
	UserID string `json:"user_id" query:"user_id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AcceptDate time.Time `json:"accept_date" query:"accept_date"`
}

// ทำเส้น check กับ save ลง tb นี้ ขึ้น
