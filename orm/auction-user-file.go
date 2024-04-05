package orm

import "github.com/FourWD/middleware/model"

type AuctionUserFile struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AuctionID            string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
	UserID               string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	FilePaymentPath      string `json:"file_payment_path" query:"file_payment_path" gorm:"type:varchar(256)"`
	FileVehicleImagePath string `json:"file_vehicle_image_path" query:"file_vehicle_image_path" gorm:"type:varchar(256)"`
}
