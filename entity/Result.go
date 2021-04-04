package entity

type Result struct {
	Id              uint32   ` json:"id"`
	FixtureId       uint32   `gorm:"not null;auto_preload" json:"fixture_id"`
	Fixture         Fixture  `gorm:"auto_preload" json:"fixture"`
	FirstClubScore  uint32   `json:"first_club_score"`
	SecondClubScore uint32   `json:"second_club_score"`
	Scorers         []Scorer `gorm:"ForeignKey:ResultId;auto_preload" json:"scorers"`
}
