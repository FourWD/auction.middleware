package orm

import "github.com/FourWD/middleware/orm"

type VehicleChecklist struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`

	FormChecklistID string `json:"form_checklist_id" query:"form_checklist_id" gorm:"type:varchar(36)"`
	IsCheck         bool   `json:"is_check" query:"is_check" gorm:"type:bool"`
}
