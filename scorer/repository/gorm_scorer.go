package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/solkn/soccer/api/entity"
)

type ScorerGormRepo struct {
	conn *gorm.DB
}

func NewClubGormRepo(db *gorm.DB) *ScorerGormRepo {
	return &ScorerGormRepo{conn: db}
}

func (scorerRepo *ScorerGormRepo) Scorers() ([]entity.Scorer, []error) {
	var scorers []entity.Scorer
	errs := scorerRepo.conn.Preload("Club").Find(&scorers).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return scorers, errs
}

func (scorerRepo *ScorerGormRepo) Scorer(id uint32) (*entity.Scorer, []error) {
	scorer := entity.Scorer{}
	errs := scorerRepo.conn.Preload("Club").First(&scorer, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &scorer, errs
}

func (scorerRepo *ScorerGormRepo) UpdateScorer(scorer *entity.Scorer) (*entity.Scorer, []error) {
	errs := scorerRepo.conn.Preload("Club").Save(scorer).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return scorer, errs
}

func (scorerRepo *ScorerGormRepo) DeleteScorer(id uint32) (*entity.Scorer, []error) {
	scorer, errs := scorerRepo.Scorer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = scorerRepo.conn.Delete(scorer, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return scorer, errs
}

func (scorerRepo *ScorerGormRepo) StoreScorer(scorer *entity.Scorer) (*entity.Scorer, []error) {
	errs := scorerRepo.conn.Preload("Club").Create(scorer).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return scorer, errs
}
