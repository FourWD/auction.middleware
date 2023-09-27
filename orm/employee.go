package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Employee struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	EmployeeTypeID string `json:"employee_type_id" query:"employee_type_id" gorm:"type:varchar(36);"`

	Username        string    `json:"username" query:"username" gorm:"type:varchar(20);"`
	Password        string    `json:"password" query:"password" gorm:"type:varchar(20);"`
	Firstname       string    `json:"firstname" query:"firstname" gorm:"type:varchar(100);"`
	Lastname        string    `json:"lastname" query:"lastname" gorm:"type:varchar(100);"`
	FileAvatarID    string    `json:"file_avartar_id" query:"file_avartar_id" gorm:"type:varchar(36)"`
	Mobile          string    `json:"mobile" query:"mobile" gorm:"type:varchar(20);"`
	Email           string    `json:"email" query:"email" gorm:"type:varchar(20);"`
	CountWrongLogin int       `json:"count_wrong_login" query:"count_wrong_login" gorm:"type:int(1);"`
	LastLoginDate   time.Time `json:"last_login_date" query:"last_login_date"`
}
