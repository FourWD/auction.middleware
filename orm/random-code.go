package orm

type RandomCode struct {
	ID   int    `json:"id" query:"id" gorm:"primaryKey;autoIncrement"`
	Code string `json:"code" query:"code" gorm:"type:varchar(5);uniqueIndex:idx_code"`
}
