package orm

type FinanceSource struct {
	FinanceID string `json:"finance_id" query:"finance_id" gorm:"type:varchar(10); uniqueIndex:idx_finance_source"`
	SourceID  string `json:"source_id" query:"source_id" gorm:"type:int(2); uniqueIndex:idx_finance_soruce"`
}
