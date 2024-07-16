package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type AuctionUserFile struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	AuctionID            string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_user_id"`
	UserID               string    `json:"user_id" query:"user_id" gorm:"type:varchar(36);uniqueIndex:idx_auction_id_user_id"`
	FilePaymentPath      string    `json:"file_payment_path" query:"file_payment_path" gorm:"type:varchar(256)"`
	FileVehicleImagePath string    `json:"file_vehicle_image_path" query:"file_vehicle_image_path" gorm:"type:varchar(256)"`
	IsGenFile            bool      `json:"is_gen_file" query:"is_gen_file" firestore:"is_gen_file" gorm:"type:bool"`
	GenFileDate          time.Time `json:"gen_file_date" query:"gen_file_date" firestore:"gen_file_date"`
	CountVehicle         int       `json:"count_vehicle" query:"count_vehicle" firestore:"count_vehicle" gorm:"type:int(6)"`
}

//พอจบรอบ auction_user_file จะหาคนชนะก่อน
