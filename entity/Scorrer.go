package entity

type Scorer struct {
	Id            uint32 ` json:"id"`
	ResultId      uint32 `json:"result_id"`
	ScorerName    string `gorm:"type:varchar(255);not null" json:"scorer_name"`
	ScoringMinute uint32 `json:"scoring_minute"`
	ClubId        uint32 `gorm:"not null;auto_preload" json:"club_id"`
	Club          Club   `gorm:"auto_preload" json:"club"`
}
