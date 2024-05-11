package model

import (
	"time"
)

type VehiclePDF struct {
	ID                  string `json:"id" `
	LicenseProvinceName string `json:"license_province_name"`

	SourceID          string `json:"source_id" query:"source_id" gorm:"type:varchar(36)"`
	SKU               string `json:"sku" query:"sku" gorm:"type:varchar(20)"`
	VehicleSubModelID string `json:"vehicle_sub_model_id" query:"vehicle_sub_model_id" gorm:"type:varchar(36)"`
	VehicleColorID    string `json:"vehicle_color_id" query:"vehicle_color_id" gorm:"type:varchar(36)"`
	ChassisNo         string `json:"chassis_no" query:"chassis_no" gorm:"type:varchar(20)"`
	EngineNo          string `json:"engine_no" query:"engine_no" gorm:"type:varchar(20)"`
	Mile              int    `json:"mile" query:"mile" gorm:"type:int(11)"`
	YearManufacturing string `json:"year_manufacturing" query:"year_manufacturing" gorm:"type:varchar(4)"`
	YearRegister      string `json:"year_register" query:"year_register" gorm:"type:varchar(4)"`
	License           string `json:"license" query:"license" gorm:"type:varchar(10)"`
	LicenseProvinceID string `json:"license_province_id" query:"license_province_id" gorm:"type:varchar(36)"`
	VehicleGradeID    string `json:"vehicle_grade_id" query:"vehicle_grade_id" gorm:"type:varchar(36)"`
	BranchID          string `json:"branch_id" query:"branch_id" gorm:"type:varchar(36)"`
	ImagePreviewPath  string `json:"image_preview_path" query:"image_preview_path" gorm:"type:varchar(400)"`
	CRPPrice          int    `json:"crp_price" query:"crp_price" gorm:"type:int"`
	Remark            string `json:"remark" query:"remark" gorm:"type:text"`
	IsRecommend       bool   `json:"is_recommend" query:"is_recommend" gorm:"type:bool"`
	//IsSpecialInterest     bool      `json:"is_special_interest" query:"is_special_interest" gorm:"type:bool"`
	EngineCapacity     string `json:"engine_capacity" query:"engine_capacity" gorm:"type:varchar(20)"`
	EngineSize         int    `json:"engine_size" query:"engine_size" gorm:"type:int"`
	EngineSizeActual   int    `json:"engine_size_actual" query:"engine_size_actual" gorm:"type:int"`
	VehicleModelID     string `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	VehicleDriveTypeID string `json:"vehicle_drive_type_id" query:"vehicle_drive_type_id" gorm:"type:varchar(2)"`
	VehicleGearID      string `json:"vehicle_gear_id" query:"vehicle_gear_id" gorm:"type:varchar(2)"`
	VehicleFuelTypeID  string `json:"vehicle_fuel_type_id" query:"vehicle_fuel_type_id" gorm:"type:varchar(2)"`
	Seat               int    `json:"seat" query:"seat" gorm:"type:int(2)"`
	// VehicleTypeID      string    `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(10)"`
	VehicleSubTypeID string `json:"vehicle_sub_type_id" query:"vehicle_sub_type_id" gorm:"type:varchar(2)"`
	// VehicleBrandID     string    `json:"vehicle_brand_id" query:"vehicle_brand_id" gorm:"type:varchar(36)"`
	RegisterTypeCode   string    `json:"register_type_code" query:"register_type_code" gorm:"type:varchar(36)"`
	RegisterType       string    `json:"register_type" query:"register_type" gorm:"type:varchar(36)"`
	LicenseReceiveDate time.Time `json:"license_receive_data" query:"license_receive_data" `
	LicenseExpireDate  time.Time `json:"license_expire_data" query:"license_data_expire" `
	TaxExpireDate      time.Time `json:"tax_expire_data" query:"license_data_expire" `
	ContractNumber     string    `json:"contract_number" query:"contract_number" gorm:"type:varchar(20)"`
	ContractDate       time.Time `json:"contract_date" query:"contract_date"`
	//VehiclePickupDate     time.Time `json:"vehicle_pickup_date" query:"vehicle_pickup_date"`
	//VehicleAuctionReceipt string    `json:"vehicle_auction_receipt" query:"vehicle_auction_receipt" gorm:"type:varchar(20)"` // เลขที่ใบรับรถขายทอดตลาด
	Accessories           string `json:"accessories" query:"accessories" gorm:"type:varchar(500)"`
	PeriodOfUse           int    `json:"period_of_use" query:"period_of_use" gorm:"type:int(2)"`
	Display               bool   `json:"display" query:"display" gorm:"type:bool"`
	AdditionalInfo        string `json:"additional_info" query:"additional_info" gorm:"type:text"`
	VehicleRegisterTypeID string `json:"vehicle_register_type_id" query:"vehicle_register_type_id" gorm:"type:varchar(2)"`
}
