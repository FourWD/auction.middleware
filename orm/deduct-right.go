package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type DeductRight struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DeductAmount        int       `json:"deduct_amount" query:"deduct_amount" gorm:"type:int"`
	DeductRight         int       `json:"deduct_right" query:"deduct_right" gorm:"type:int"`
	UserID              string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	DeductDate          time.Time `json:"deduct_date" query:"deduct_date"`
	DeductRightReasonID string    `json:"deduct_right_reason_id" query:"deduct_right_reason_id" gorm:"type:varchar(2)"`
	IsApprove           bool      `json:"is_approve" query:"is_approve" gorm:"type:bool"`
	ApproveBy           string    `json:"approve_by" query:"approve_by" gorm:"type:varchar(36)"`
	ApproveDate         time.Time `json:"approve_date" query:"approve_date"`
	EffectiveDate       time.Time `json:"effective_date" query:"effective_date"`
	Remark              string    `json:"remark" query:"remark" gorm:"type:text;"`
}
