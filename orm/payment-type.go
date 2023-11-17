package orm

import "github.com/FourWD/middleware/orm"

type PaymentType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name           string `json:"name" query:"name" gorm:"type:varchar(50);"`
	PaymentGroupID string `json:"payment_group_id" query:"payment_group_id" gorm:"type:varchar(10);"`
	orm.GormRowOrder
}

/*
01 เงินสด PAYMENT
02 ขอสินเชื่อ PAYMENT
03 โอนเงิน REGISTER
04 โมบายแบ๊งค์กิ้งค์ REGISTER
*/
