package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type UserAccpetVat struct {
	UserID string `json:"user_id" query:"user_id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	AcceptDate time.Time `json:"accept_date" query:"accept_date"`
}

// ทำเส้น check กับ save ลง tb นี้ ขึ้น
