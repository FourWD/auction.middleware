package orm

type RoundSource struct {
	RoundID  string `json:"round_id" query:"round_id" gorm:"type:varchar(36); uniqueIndex:idx_round_source"`
	SourceID string `json:"source_id" query:"source_id" gorm:"type:varchar(10); uniqueIndex:idx_round_source"`
}
