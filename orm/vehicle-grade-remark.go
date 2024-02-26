package orm

type VehicleGradeRemark struct {
	ChassisNumber string `json:"chassis_number" query:"chassis_number" gorm:"type:varchar(20)"`
	GradeRemark   string `json:"grade_remark" query:"grade_remark" gorm:"type:text"`
}
