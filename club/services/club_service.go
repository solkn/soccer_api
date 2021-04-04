package services

import (
	club "github.com/solkn/soccer/api/club"
	"github.com/solkn/soccer/api/entity"
)

type ClubService struct {
	clubRepo club.ClubsRepository
}

func NewClubService(clubRepo club.ClubsRepository) *ClubService {
	return &ClubService{clubRepo: clubRepo}
}

func (cs *ClubService) Clubs() ([]entity.Club, []error) {

	clubs, errs := cs.clubRepo.Clubs()
	if len(errs) > 0 {
		return nil, errs
	}
	return clubs, errs

}

func (cs *ClubService) Club(id uint32) (*entity.Club, []error) {
	cl, errs := cs.clubRepo.Club(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cl, errs

}

func (cs *ClubService) UpdateClub(club *entity.Club) (*entity.Club, []error) {
	cl, errs := cs.clubRepo.UpdateClub(club)
	if len(errs) > 0 {
		return nil, errs
	}
	return cl, errs

}

func (cs *ClubService) DeleteClub(id uint32) (*entity.Club, []error) {

	cl, errs := cs.clubRepo.DeleteClub(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cl, errs
}

func (cs *ClubService) StoreClub(club *entity.Club) (*entity.Club, []error) {

	cl, errs := cs.clubRepo.StoreClub(club)
	if len(errs) > 0 {
		return nil, errs
	}
	return cl, errs
}
