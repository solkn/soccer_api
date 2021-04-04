package services

import (
	"github.com/solkn/soccer/api/entity"
	scorer "github.com/solkn/soccer/api/scorer"
)

type ScorerService struct {
	scorerRepo scorer.ScorersRepository
}

func NewScorerService(scorerRepo scorer.ScorersRepository) *ScorerService {
	return &ScorerService{scorerRepo: scorerRepo}
}

func (ss *ScorerService) Scorers() ([]entity.Scorer, []error) {

	scorers, errs := ss.scorerRepo.Scorers()
	if len(errs) > 0 {
		return nil, errs
	}
	return scorers, errs

}

func (ss *ScorerService) Scorer(id uint32) (*entity.Scorer, []error) {
	scrr, errs := ss.scorerRepo.Scorer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return scrr, errs

}

func (ss *ScorerService) UpdateScorer(scorer *entity.Scorer) (*entity.Scorer, []error) {
	scrr, errs := ss.scorerRepo.UpdateScorer(scorer)
	if len(errs) > 0 {
		return nil, errs
	}
	return scrr, errs

}

func (ss *ScorerService) DeleteScorer(id uint32) (*entity.Scorer, []error) {

	scrr, errs := ss.scorerRepo.DeleteScorer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return scrr, errs
}

func (ss *ScorerService) StoreScorer(scorer *entity.Scorer) (*entity.Scorer, []error) {

	scrr, errs := ss.scorerRepo.StoreScorer(scorer)
	if len(errs) > 0 {
		return nil, errs
	}
	return scrr, errs
}
