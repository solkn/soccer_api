package services

import (
	"github.com/solkn/soccer/api/entity"
	result "github.com/solkn/soccer/api/result"
)

type ResultService struct {
	resultsRepo result.ResultsRepository
}

func NewClubService(resultRepo result.ResultsRepository) *ResultService {
	return &ResultService{resultsRepo: resultRepo}
}

func (rs *ResultService) Results() ([]entity.Result, []error) {

	results, errs := rs.resultsRepo.Results()
	if len(errs) > 0 {
		return nil, errs
	}
	return results, errs

}

func (rs *ResultService) Result(id uint32) (*entity.Result, []error) {
	rst, errs := rs.resultsRepo.Result(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rst, errs

}

func (rs *ResultService) UpdateResult(result *entity.Result) (*entity.Result, []error) {
	rslt, errs := rs.resultsRepo.UpdateResult(result)
	if len(errs) > 0 {
		return nil, errs
	}
	return rslt, errs

}

func (rs *ResultService) DeleteResult(id uint32) (*entity.Result, []error) {

	rslt, errs := rs.resultsRepo.DeleteResult(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rslt, errs
}

func (rs *ResultService) StoreResult(result *entity.Result) (*entity.Result, []error) {

	rslt, errs := rs.resultsRepo.StoreResult(result)
	if len(errs) > 0 {
		return nil, errs
	}
	return rslt, errs
}
