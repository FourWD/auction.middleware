package orm

import (
	"github.com/FourWD/middleware/model"
)

type TemplateDowloadVehicleImage struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	VehicleImageID string  `json:"vehicle_image_id" query:"vehicle_image_id" gorm:"type:varchar(36)"`
	RowOrder       float32 `json:"row_order" query:"row_order" gorm:"type:decimal(5,2)"`
}
