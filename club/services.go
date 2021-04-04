package club

import "github.com/solkn/soccer/api/entity"

type ClubsServices interface {
	Clubs() ([]entity.Club, []error)
	Club(id uint32) (*entity.Club, []error)
	StoreClub(user *entity.Club) (*entity.Club, []error)
	UpdateClub(order *entity.Club) (*entity.Club, []error)
	DeleteClub(id uint32) (*entity.Club, []error)
}
