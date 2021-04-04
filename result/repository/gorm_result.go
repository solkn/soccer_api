package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/solkn/soccer/api/entity"
)

type ResultGormRepo struct {
	conn *gorm.DB
}

func NewResultGormRepo(db *gorm.DB) *ResultGormRepo {
	return &ResultGormRepo{conn: db}
}

func (resultRepo *ResultGormRepo) Results() ([]entity.Result, []error) {
	var results []entity.Result
	errs := resultRepo.conn.Preload("Fixture").Preload("Fixture.Clubs").Preload("Scorers").Preload("Scorers.Club").Find(&results).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return results, errs
}

func (resultRepo *ResultGormRepo) Result(id uint32) (*entity.Result, []error) {
	result := entity.Result{}
	errs := resultRepo.conn.Preload("Fixture").Preload("Fixture.Clubs").Preload("Scorers").Preload("Scorers.Club").First(&result, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &result, errs
}

func (resultRepo *ResultGormRepo) UpdateResult(result *entity.Result) (*entity.Result, []error) {
	errs := resultRepo.conn.Preload("Fixture").Preload("Fixture.Clubs").Preload("Scorers").Preload("Scorers.Club").Save(result).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return result, errs
}

func (resultRepo *ResultGormRepo) DeleteResult(id uint32) (*entity.Result, []error) {
	result, errs := resultRepo.Result(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = resultRepo.conn.Delete(result, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return result, errs
}

func (resultRepo *ResultGormRepo) StoreResult(result *entity.Result) (*entity.Result, []error) {
	errs := resultRepo.conn.Preload("Fixture").Preload("Fixture.Clubs").Preload("Scorers").Preload("Scorers.Club").Create(result).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return result, errs
}
