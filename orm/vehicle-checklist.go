package orm

import "github.com/FourWD/middleware/orm"

type VehicleChecklist struct { // 01 Seller 02 Customer
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleID string `json:"vehicle_id" query:"vehicle_id" db:"vehicle_id" gorm:"type:varchar(36)"`

	FileChecklistID string `json:"file_checklist_id" query:"file_checklist_id" db:"file_checklist_id" gorm:"type:varchar(36)"`
}
