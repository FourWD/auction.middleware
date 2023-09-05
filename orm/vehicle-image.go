package orm

import "github.com/FourWD/middleware/orm"

type VehicleImage struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleID              string `json:"vehicle_id" query:"vehicle_id" db:"vehicle_id" gorm:"type:varchar(36)"`
	TemplateVehicleImageID string `json:"template_vehicle_image_id" query:"template_vehicle_image_id" gorm:"type:varchar(36);"`

	ImagePath string `json:"image_path" query:"image_path" gorm:"type:varchar(255)"`
	IsDelete  bool   `json:"is_delete" query:"is_delete" gorm:"type:bool"`

	orm.RowOrder
}