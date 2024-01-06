package orm

type LogProxyType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(1);primary_key"`

	Name string `json:"name" query:"name" gorm:"type:varchar(50)"`
}
