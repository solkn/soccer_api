package entity

import "time"

type Fixture struct {
	Id               uint32    `json:"id"`
	StartingDate     time.Time `gorm:"default:current_timestamp" json:"starting_date"`
	Clubs            []Club    `gorm:"many2many:fixture_clubs;auto_preload;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"clubs"`
	StadiumLatitude  float64   `json:"stadium_latitude"`
	StadiumLongitude float64   `json:"stadium_longitude"`
	RefereeName      string    `gorm:"type:varchar(255)" json:"referee_name"`
}
