package orm

type RandomCode struct {
	ID   int    `json:"id" query:"id" gorm:"primaryKey"`
	Code string `json:"code" query:"code" gorm:"type:varchar(4);uniqueIndex:idx_code"`
}
