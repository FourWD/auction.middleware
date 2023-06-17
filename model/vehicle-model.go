package model

import (
	"github.com/FourWD/middleware/orm"
)

type VehicleModel struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleModelID string `json:"vehicle_model_id" query:"vehicle_model_id" db:"vehicle_model_id" gorm:"type:varchar(36);"`

	VehicleColorID    string `json:"vehicle_color_id" query:"vehicle_color_id" db:"vehicle_color_id" gorm:"type:varchar(36);"`
	ChassisNo         string `json:"chassis_no" query:"chassis_no" db:"chassis_no" gorm:"type:varchar(20);"`
	Mile              string `json:"mile" query:"mile" db:"mile" gorm:"type:int(11);"`
	YearManufacturing string `json:"year_manufacturing" query:"year_manufacturing" db:"year_manufacturing" gorm:"type:int(4);"`
	YearRegister      string `json:"year_register" query:"year_register" db:"year_register" gorm:"type:int(4);"`

	License          string `json:"license" query:"license" db:"license" gorm:"type:varchar(20);"`
	LicenseProviceID string `json:"license_province_id" query:"license_province_id" db:"license_province_id" gorm:"type:varchar(36);"`
	VehicleGradeID   string `json:"vehicle_grade_id" query:"vehicle_grade_id" db:"vehicle_grade_id" gorm:"type:varchar(36);"`
	OpenPrice        string `json:"open_price" query:"open_price" db:"open_price" gorm:"type:varchar(20);"`
	MinPrice         string `json:"min_price" query:"min_price" db:"min_price" gorm:"type:varchar(20);"`
	BranchID         string `json:"branch_id" query:"branch_id" db:"branch_id" gorm:"type:varchar(36)"`
	Remark           string `json:"remark" query:"remark" db:"name" gorm:"type:text;"`
}
