package orm

import "github.com/FourWD/middleware/orm"

type Favorite struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FavoriteTypeID string `json:"favorite_type_id" query:"favorite_type_id" gorm:"type:varchar(36)"`
	UserID         string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	ItemID         string `json:"item_id" query:"item_id" gorm:"type:varchar(36)"`
	ItemTableName  string `json:"item_table_name" query:"item_table_name" gorm:"type:varchar(50)"`
}
