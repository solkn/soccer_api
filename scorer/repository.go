package scorer

import "github.com/solkn/soccer/api/entity"

type ScorersRepository interface {
	Scorers() ([]entity.Scorer, []error)
	Scorer(id uint32) (*entity.Scorer, []error)
	StoreScorer(user *entity.Scorer) (*entity.Scorer, []error)
	UpdateScorer(order *entity.Scorer) (*entity.Scorer, []error)
	DeleteScorer(id uint32) (*entity.Scorer, []error)
}
