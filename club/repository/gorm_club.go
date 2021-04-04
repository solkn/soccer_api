package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/solkn/soccer/api/entity"
)

type ClubGormRepo struct {
	conn *gorm.DB
}

func NewClubGormRepo(db *gorm.DB) *ClubGormRepo {
	return &ClubGormRepo{conn: db}
}

func (clubRepo *ClubGormRepo) Clubs() ([]entity.Club, []error) {
	var clubs []entity.Club
	errs := clubRepo.conn.Find(&clubs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return clubs, errs
}

func (clubRepo *ClubGormRepo) Club(id uint32) (*entity.Club, []error) {
	club := entity.Club{}
	errs := clubRepo.conn.First(&club, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &club, errs
}

func (clubRepo *ClubGormRepo) UpdateClub(club *entity.Club) (*entity.Club, []error) {
	errs := clubRepo.conn.Save(club).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return club, errs
}

func (clubRepo *ClubGormRepo) DeleteClub(id uint32) (*entity.Club, []error) {
	club, errs := clubRepo.Club(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = clubRepo.conn.Delete(club, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return club, errs
}

func (clubRepo *ClubGormRepo) StoreClub(club *entity.Club) (*entity.Club, []error) {
	errs := clubRepo.conn.Create(club).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return club, errs
}
