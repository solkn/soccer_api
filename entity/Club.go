package entity

type Club struct {
	Id   uint32 `json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"club_name"`
}
