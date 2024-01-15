package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type Branch struct {
	midOrm.Branch

	IsShowWatermask bool `json:"is_show_watermask" query:"is_show_watermask" gorm:"type:bool"`
}
