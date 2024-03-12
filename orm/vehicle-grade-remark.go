package orm

import "time"

type VehicleGradeRemark struct {
	ChassisNumber string    `json:"chassis_number" query:"chassis_number" gorm:"type:varchar(50);primary_key"`
	License       string    `json:"license" query:"license" gorm:"type:varchar(50)"`
	GradeRemark   string    `json:"grade_remark" query:"grade_remark" gorm:"type:text"`
	Grade         string    `json:"grade" query:"grade" gorm:"type:varchar(2)"`
	UpdatedAt     time.Time `json:"updated_at" query:"updated_at"`
}
