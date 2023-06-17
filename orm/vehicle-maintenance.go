package orm

import "github.com/FourWD/middleware/orm"

type VehicleMaintenance struct { // 01 Seller 02 Customer
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleID string `json:"vehicle_id" query:"vehicle_id" db:"vehicle_id" gorm:"type:varchar(36)"`

	MaintenancedAt        string `json:"maintenanced_at" query:"maintenanced_at" db:"maintenanced_at" gorm:"type:varchar(20);"`
	Mile                  int    `json:"mile" query:"mile" db:"mile" gorm:"type:int(11);"`
	MaintenanceLocationID string `json:"maintenanced_location_id" query:"maintenanced_location_id" db:"maintenanced_location_id" gorm:"type:varchar(36);"`
	MaintenanceTypeID     string `json:"maintenanced_type_id" query:"maintenanced_type_id" db:"maintenanced_type_id" gorm:"type:varchar(36);"`
	Remark                string `json:"remark" query:"remark" db:"name" gorm:"type:varchar(500);"`
}
