package model

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type VehicleModel struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SourceID               string    `json:"source_id" query:"source_id" gorm:"type:varchar(36);"`
	VehicleSubmodelID      string    `json:"vehicle_submodel_id" query:"vehicle_submodel_id" gorm:"type:varchar(36);"`
	VehicleColorID         string    `json:"vehicle_color_id" query:"vehicle_color_id" gorm:"type:varchar(36);"`
	ChassisNo              string    `json:"chassis_no" query:"chassis_no" gorm:"type:varchar(20);"`
	EngineNo               string    `json:"engine_no" query:"engine_no" gorm:"type:varchar(20);"`
	Mile                   int       `json:"mile" query:"mile" gorm:"type:int(11);"`
	YearManufacturing      int       `json:"year_manufacturing" query:"year_manufacturing" gorm:"type:int(4);"`
	YearRegister           int       `json:"year_register" query:"year_register" gorm:"type:int(4);"`
	License                string    `json:"license" query:"license" gorm:"type:varchar(10);"`
	LicenseProviceID       string    `json:"license_province_id" query:"license_province_id" gorm:"type:varchar(36);"`
	VehicleGradeID         string    `json:"vehicle_grade_id" query:"vehicle_grade_id" gorm:"type:varchar(36);"`
	BranchID               string    `json:"branch_id" query:"branch_id" gorm:"type:varchar(36)"`
	ImagePreviewPath       string    `json:"image_preview_path" query:"image_preview_path" gorm:"type:varchar(400)"`
	CRPPrice               int       `json:"crp_price" query:"crp_price" gorm:"type:int"`
	Remark                 string    `json:"remark" query:"remark" gorm:"type:text;"`
	SoldStatusID           int       `json:"sold_status_id" query:"sold_status_id" gorm:"type:int"`
	DisplayRecommend       bool      `json:"display_recommend" query:"display_recommend" gorm:"type:bool"`
	DisplaySpecialInterest bool      `json:"display_special_interest" query:"display_special_interest" gorm:"type:bool"`
	EngineCapacity         int       `json:"engine_capacity" query:"engine_capacity" gorm:"type:int"`
	EngineSize             int       `json:"engine_size" query:"engine_size" gorm:"type:int"`
	EngineSizeActual       int       `json:"engine_size_actual" query:"engine_size_actual" gorm:"type:int"`
	VehicleModelID         string    `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36);"`
	VehicleDriveTypeID     string    `json:"vehicle_drive_type_id" query:"vehicle_drive_type_id" gorm:"type:varchar(2);"`
	VehicleGearID          string    `json:"vehicle_gear_id" query:"vehicle_gear_id" gorm:"type:varchar(2);"`
	VehicleFuelID          string    `json:"vehicle_fuel_id" query:"vehicle_fuel_id" gorm:"type:varchar(2);"`
	Seat                   int       `json:"seat" query:"seat" gorm:"type:int(2);"`
	VehicleTypeID          string    `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(36);"`
	VehicleBrandID         string    `json:"vehicle_brand_id" query:"vehicle_brand_id" gorm:"type:varchar(36);"`
	RegisterTypeCode       string    `json:"register_type_code" query:"register_type_code" gorm:"type:varchar(36);"`
	RegisterType           string    `json:"register_type" query:"register_type" gorm:"type:varchar(36);"`
	LicenseReceiveDate     string    `json:"license_receive_data" query:"license_receive_data" gorm:"type:varchar(20);"`
	LicenseExpireDate      string    `json:"license_expire_data" query:"license_data_expire" gorm:"type:varchar(20);"`
	TaxExpireDate          string    `json:"tax_expire_data" query:"license_data_expire" gorm:"type:varchar(20);"`
	ContractNumber         string    `json:"contract_number" query:"contract_number" gorm:"type:varchar(20);"`
	VehiclePickupDate      time.Time `json:"vehicle_pickup_date" query:"vehicle_pickup_date"`
	VehicleAuctionReceipt  string    `json:"vehicle_auction_receipt" query:"vehicle_auction_receipt" gorm:"type:varchar(20);"` // เลขที่ใบรับรถขายทอดตลาด
	Accessories            string    `json:"accessories" query:"accessories" gorm:"type:varchar(500);"`
	PeriodOfUse            int       `json:"period_of_use" query:"period_of_use" gorm:"type:int(2);"`
	Display                bool      `json:"display" query:"display" gorm:"type:bool"`
	AdditionalInfo         string    `json:"additional_info" query:"additional_info" gorm:"type:text;"`
	A3FirstName            string    `json:"a3_first_name" query:"a3_first_name" gorm:"type:varchar(50);"`
	A3LastName             string    `json:"a3_last_name" query:"a3_last_name" gorm:"type:varchar(50);"`
	A3CarOwnerNumber       int       `json:"a3_car_owner_number" query:"a3_car_owner_number" gorm:"type:int(2);"`
	A3AccidentHistory      string    `json:"a3_accident_history" query:"a3_accident_history" gorm:"type:varchar(50);"`
	A3VehicleLienExists    string    `json:"a3_vehicle_lien_exists" query:"a3_vehicle_lien_exists" gorm:"type:varchar(50);"`

	// ImgStrFront        string `json:"img_str_front" query:"img_str_front" gorm:"type:varchar(400)"`
	// ImgStrBack         string `json:"img_str_back" query:"img_str_back" gorm:"type:varchar(400)"`
	// ImgStrRight        string `json:"img_str_right" query:"img_str_right" gorm:"type:varchar(400)"`
	// ImgStrLeft         string `json:"img_str_left" query:"img_str_left" gorm:"type:varchar(400)"`
	// ImgFrontLeft45     string `json:"img_front_left_45" query:"img_front_left_45" gorm:"column:img_front_left_45;type:varchar(400)"`
	// ImgFrontRight45    string `json:"img_front_right_45" query:"img_front_right_45" gorm:"column:img_front_right_45;type:varchar(400)"`
	// ImgBackLeft45      string `json:"img_back_left_45" query:"img_back_left_45" gorm:"column:img_back_left_45;type:varchar(400)"`
	// ImgBackRight45     string `json:"img_back_right_45" query:"img_back_right_45" gorm:"column:img_back_right_45;type:varchar(400)"`
	// ImgInFront         string `json:"img_in_front" query:"img_in_front" gorm:"type:varchar(400)"`
	// ImgInBack          string `json:"img_in_back" query:"img_in_back" gorm:"type:varchar(400)"`
	// ImgConsole         string `json:"img_console" query:"img_console" gorm:"type:varchar(400)"`
	// ImgMileage         string `json:"img_mileage" query:"img_mileage" gorm:"type:varchar(400)"`
	// ImgVehTools        string `json:"img_veh_tools" query:"img_veh_tools" gorm:"type:varchar(400)"`
	// ImgEngineRoom      string `json:"img_engine_room" query:"img_engine_room" gorm:"type:varchar(400)"`
	// ImgGas             string `json:"img_gas" query:"img_gas" gorm:"type:varchar(400)"`
	// ImgOut360          string `json:"img_out_360" query:"img_out_360" gorm:"column:img_out_360;type:varchar(400)"`
	// ImgIn360           string `json:"img_in_360" query:"img_in_360" gorm:"column:img_in_360;type:varchar(400)"`
	// ImgAct             string `json:"img_act" query:"img_act" gorm:"type:varchar(400)"`
	// ImgInsurance       string `json:"img_insurance" query:"img_insurance" gorm:"type:varchar(400)"`
	// ImgInspectionFront string `json:"img_inspection_front" query:"img_inspection_front" gorm:"type:varchar(400)"`
	// ImgInspectionBack  string `json:"img_inspection_back" query:"img_inspection_back" gorm:"type:varchar(400)"`
}
