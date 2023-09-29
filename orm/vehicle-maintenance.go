package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type VehicleMaintenance struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleID             string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	MaintenanceDate       time.Time `json:"maintenance_date" query:"maintenance_date"`
	Mile                  int       `json:"mile" query:"mile" gorm:"type:int(11);"`
	MaintenanceLocationID string    `json:"maintenanced_location_id" query:"maintenanced_location_id" gorm:"type:varchar(36);"`
	MaintenanceTypeID     string    `json:"maintenanced_type_id" query:"maintenanced_type_id" gorm:"type:varchar(36);"`
	Remark                string    `json:"remark" query:"remark" gorm:"type:text;"`
}
