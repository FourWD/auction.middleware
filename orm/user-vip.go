package orm

import "github.com/FourWD/middleware/orm"

type UserVip struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID        string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	UserVipTypeID string `json:"user_vip_type_id" query:"user_vip_type_id" gorm:"type:varchar(2)"`

	StartDate string `json:"start_date" query:"start_date" gorm:"not null;type:varchar(30)"`
	EndDate   string `json:"end_date" query:"end_date" gorm:"not null;type:varchar(30)"`
	VipRemark string `json:"vip_remark" query:"vip_remark"  gorm:"type:text"`
	VipBy     string `json:"vip_by" query:"vip_by" gorm:"type:varchar(36)"` // ban approve

	IsCancel       bool   `json:"is_cancle" query:"is_cancle" gorm:"bool"`
	ExpireDatetime string `json:"expire_datetime" query:"expire_datetime" gorm:"not null;type:varchar(30)"`
	CancelBy       string `json:"cancel_by" query:"cancel_by" gorm:"type:varchar(36)"`
	CancelRemark   string `json:"cancel_remark" query:"cancel_remark" gorm:"type:text"`
}
