package orm

import "github.com/FourWD/middleware/orm"

type FavoriteVehicle struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID           string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	AuctionID        string `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
	VehicleID        string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	VehicleTableName string `json:"vehicle_table_name" query:"vehicle_table_name" gorm:"type:varchar(50)"`
	IsDelete         string `json:"is_delete" query:"is_delete" gorm:"type:bool"`
}
