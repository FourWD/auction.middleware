package model

import (
	"github.com/FourWD/middleware/orm"
)

type VehicleModel struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SourceID          string `json:"source_id" query:"source_id" gorm:"type:varchar(36);"`
	VehicleModelID    string `json:"vehicle_model_id" query:"vehicle_model_id" db:"vehicle_model_id" gorm:"type:varchar(36);"`
	VehicleSubmodelID string `json:"vehicle_submodel_id" query:"vehicle_submodel_id" gorm:"type:varchar(36);"`

	VehicleColorID    string `json:"vehicle_color_id" query:"vehicle_color_id" db:"vehicle_color_id" gorm:"type:varchar(36);"`
	ChassisNo         string `json:"chassis_no" query:"chassis_no" db:"chassis_no" gorm:"type:varchar(20);"`
	Mile              string `json:"mile" query:"mile" db:"mile" gorm:"type:int(11);"`
	YearManufacturing string `json:"year_manufacturing" query:"year_manufacturing" db:"year_manufacturing" gorm:"type:int(4);"`
	YearRegister      string `json:"year_register" query:"year_register" db:"year_register" gorm:"type:int(4);"`

	VehicleNo string `json:"vehicle_no" query:"vehicle_no" db:"vehicle_no" gorm:"type:varchar(10);"`

	License          string `json:"license" query:"license" db:"license" gorm:"type:varchar(10);"`
	LicenseProviceID string `json:"license_province_id" query:"license_province_id" db:"license_province_id" gorm:"type:varchar(36);"`
	VehicleGradeID   string `json:"vehicle_grade_id" query:"vehicle_grade_id" db:"vehicle_grade_id" gorm:"type:varchar(36);"`
	BranchID         string `json:"branch_id" query:"branch_id" db:"branch_id" gorm:"type:varchar(36)"`
	Remark           string `json:"remark" query:"remark" db:"name" gorm:"type:text;"`

	ImgStrFront        string `json:"img_str_front" query:"img_str_front" db:"img_str_front" gorm:"type:varchar(400)"`
	ImgStrBack         string `json:"img_str_back" query:"img_str_back" db:"img_str_back" gorm:"type:varchar(400)"`
	ImgStrRight        string `json:"img_str_right" query:"img_str_right" db:"img_str_right" gorm:"type:varchar(400)"`
	ImgStrLeft         string `json:"img_str_left" query:"img_str_left" db:"img_str_left" gorm:"type:varchar(400)"`
	ImgFrontLeft45     string `json:"img_front_left_45" query:"img_front_left_45" db:"img_front_left_45" gorm:"column:img_front_left_45;type:varchar(400)"`
	ImgFrontRight45    string `json:"img_front_right_45" query:"img_front_right_45" db:"img_front_right_45" gorm:"column:img_front_right_45;type:varchar(400)"`
	ImgBackLeft45      string `json:"img_back_left_45" query:"img_back_left_45" db:"img_back_left_45" gorm:"column:img_back_left_45;type:varchar(400)"`
	ImgBackRight45     string `json:"img_back_right_45" query:"img_back_right_45" db:"img_back_right_45" gorm:"column:img_back_right_45;type:varchar(400)"`
	ImgInFront         string `json:"img_in_front" query:"img_in_front" db:"img_in_front" gorm:"type:varchar(400)"`
	ImgInBack          string `json:"img_in_back" query:"img_in_back" db:"img_in_back" gorm:"type:varchar(400)"`
	ImgConsole         string `json:"img_console" query:"img_console" db:"img_console" gorm:"type:varchar(400)"`
	ImgMileage         string `json:"img_mileage" query:"img_mileage" db:"img_mileage" gorm:"type:varchar(400)"`
	ImgVehTools        string `json:"img_veh_tools" query:"img_veh_tools" db:"img_veh_tools" gorm:"type:varchar(400)"`
	ImgEngineRoom      string `json:"img_engine_room" query:"img_engine_room" db:"img_engine_room" gorm:"type:varchar(400)"`
	ImgGas             string `json:"img_gas" query:"img_gas" db:"img_gas" gorm:"type:varchar(400)"`
	ImgOut360          string `json:"img_out_360" query:"img_out_360" db:"img_out_360" gorm:"column:img_out_360;type:varchar(400)"`
	ImgIn360           string `json:"img_in_360" query:"img_in_360" db:"img_in_360" gorm:"column:img_in_360;type:varchar(400)"`
	ImgAct             string `json:"img_act" query:"img_act" db:"img_act" gorm:"type:varchar(400)"`
	ImgInsurance       string `json:"img_insurance" query:"img_insurance" db:"img_insurance" gorm:"type:varchar(400)"`
	ImgInspectionFront string `json:"img_inspection_front" query:"img_inspection_front" db:"img_inspection_front" gorm:"type:varchar(400)"`
	ImgInspectionBack  string `json:"img_inspection_back" query:"img_inspection_back" db:"img_inspection_back" gorm:"type:varchar(400)"`
}
