package orm

import "github.com/FourWD/middleware/orm"

type TemplateVehicleImage struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleImageGroupID string `json:"vehicle_type_id" query:"vehicle_type_id" db:"vehicle_type_id" gorm:"type:varchar(36);"`

	Name string `json:"name" query:"name" db:"name" gorm:"type:varchar(20);"`

	orm.RowOrder
}

// รูป ด้านซ้าย ด้านขวา ของภายหน้า
