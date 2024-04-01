package utils

import (
	"time"

	"github.com/FourWD/auction.middleware/orm"
)

type VehicleImageD struct {
	Vehicle orm.Vehicle
	Images  []Image `json:"images"`
}
type UserSummary struct {
	// user orm.User
	Code       string `json:"code"`
	UserTypeID string `json:"user_type_id"`
	// FirstName string `json:"firstname"`
	Mobile          string `json:"mobile"`
	UserFirstname   string `json:"user_firstname" query:"user_firstname"`
	UserLastname    string `json:"user_lastname" query:"user_lastname"`
	CompanyName     string `json:"company_name"`
	CompanyPhone    string `json:"company_phone"`
	Street          string `json:"street"`
	Address         string `json:"address"`
	Post            string `json:"post"`
	Email           string `json:"email"`
	DistrictName    string `json:"district_name"`
	SubDistrictName string `json:"sub_district_name"`
	ProvinceName    string `json:"province_name"`
	ProvinceID      string `json:"province_id"`

	AuctionName string `json:"auction_name"`

	//tax
	IsAddressTax int `json:"is_address_tax"`

	TaxProvinceName    string `json:"tax_province_name"`
	TaxDistrictName    string `json:"tax_district_name"`
	TaxProvinceID      string `json:"tax_province_id"`
	TaxAddress         string `json:"tax_address"`
	TaxBuilding        string `json:"tax_building"`
	TaxRoom            string `json:"tax_room"`
	TaxStreet          string `json:"tax_street"`
	TaxSubDistrictName string `json:"tax_sub_district_name"`
	TaxPostCode        string `json:"tax_post_code"`
	TaxBranchName      string `json:"tax_branch_name"`
}
type VehicleSummary struct {
	CodeAuction          string `json:"code_auction"`
	ClosePrice           string `json:"close_price"`
	CarImage             string `json:"car_image"`
	LicenseImage         string `json:"license_image"`
	VehicleAuctionRecipt string `json:"vehicle_auction_receipt"`
	VehicleGradeID       string `json:"vehicle_grade_id"`
	VehicleNo            string `json:"vehicle_no"`
	TotalCar             string `json:"total_car"`
	VehicleBrandName     string `json:"vehicle_brand_name"`
	VehicleSubModelName  string `json:"vehicle_model_name"`
	VehicleModelName     string `json:"vehicle_sub_model_name"`
	VehicleColorName     string `json:"vehicle_color_name"`
	YearManuFacturing    string `json:"year_manufacturing"`
	YearRegister         string `json:"year_register"`
	ImagePreviewPath     string `json:"image_preview_path"`
	OpenPrice            string `json:"open_price"`
	AuctionName          string `json:"auction_name"`
	TaxExprieDate        string `json:"tax_expire_date"`
	Mile                 string `json:"mile"`
	ChassisNo            string `json:"chassis_no"`
	EngineSize           string `json:"engine_size"`
	EngineNo             string `json:"engine_no"`

	EndDate     time.Time `json:"end_date"`
	Remark      string    `json:"remark"`
	License     string    `json:"license"`
	BranchLabel string    `json:"branch_label"`
	CarPrice10  string    `json:"car_price_10"`
	CarPrice90  string    `json:"car_price_90"`

	LicenseProvinceName string `json:"license_province_name"`
}
type Summary struct {
	RoundName   string `json:"round_name"`
	AuctionName string `json:"auction_name"`
	TotalCar    int    `json:"total_car"`
	TotalPrice  int    `json:"total_price"`
}

//vehicle carlist

type AuctionVehicleDetail struct {
	// Auction     string
	Auction  []AuctionList
	Vehicles []VehicleCarlist
}
type AuctionList struct {
	// Auction     string

	AuctionName string `json:"auction_name"`
	RoundName   string `json:"round_name"`
	Name        string `json:"name"`
}

type VehicleCarlist struct {
	// Vehicle              []orm.Vehicle
	VehicleAuctionRecipt string `json:"vehicle_auction_receipt"`
	VehicleGradeID       string `json:"vehicle_grade_id"`
	VehicleGradeName     string `json:"vehicle_grade_name"`

	TotalCar            string `json:"total_car"`
	ID                  string `json:"id"`
	VehicleBrandName    string `json:"vehicle_brand_name"`
	VehicleSubModelName string `json:"vehicle_model_name"`
	VehicleModelName    string `json:"vehicle_sub_model_name"`
	VehicleColorName    string `json:"vehicle_color_name"`
	VehicleNo           string `json:"vehicle_no"`

	Year              string    `json:"year"`
	ImagePreviewPath  string    `json:"image_preview_path"`
	OpenPrice         string    `json:"open_price"`
	TaxExpireDate     time.Time `json:"tax_expire_date" query:"tax_expire_date"`
	LicenseExpireDate string    `json:"license_expire_date" query:"license_expire_date"`

	Mile        string `json:"mile"`
	Remark      string `json:"remark"`
	License     string `json:"license"`
	Label       string `json:"label"`
	ChassisNo   string `json:"chassis_no"`
	GradeRemark string `json:"grade_remark"`

	LicenseProvinceName string `json:"license_province_name"`
}
