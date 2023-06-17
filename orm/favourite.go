package orm

import "github.com/FourWD/middleware/orm"

type Favourite struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FavouriteTypeID string `json:"favourite_type_id" query:"favourite_type_id" db:"favourite_type_id" gorm:"type:varchar(36)"`
	UserID          string `json:"user_id" query:"user_id" db:"user_id" gorm:"type:varchar(36)"`
	ItemID          string `json:"item_id" query:"item_id" db:"item_id" gorm:"type:varchar(36)"`
}
