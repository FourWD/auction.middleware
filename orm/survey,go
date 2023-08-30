package orm

import "github.com/FourWD/middleware/orm"

type Survey struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SurveyGroupID string `json:"survey_group_id" query:"survey_group_id" gorm:"type:varchar(36)"`
	Subject       string `json:"subject" query:"subject" gorm:"type:varchar(100)"`
	AnswerCount   int    `json:"answer_count" query:"answer_count" gorm:"type:int"`
	orm.RowOrder
}
