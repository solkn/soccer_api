package result

import "github.com/solkn/soccer/api/entity"

type ResultsRepository interface {
	Results() ([]entity.Result, []error)
	Result(id uint32) (*entity.Result, []error)
	StoreResult(user *entity.Result) (*entity.Result, []error)
	UpdateResult(order *entity.Result) (*entity.Result, []error)
	DeleteResult(id uint32) (*entity.Result, []error)
}
