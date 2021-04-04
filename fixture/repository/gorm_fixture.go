package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/solkn/soccer/api/entity"
)

type FixtureGormRepo struct {
	conn *gorm.DB
}

func NewFixtureGormRepo(db *gorm.DB) *FixtureGormRepo {
	return &FixtureGormRepo{conn: db}
}

func (fixtureRepo *FixtureGormRepo) Fixtures() ([]entity.Fixture, []error) {
	var fixtures []entity.Fixture
	errs := fixtureRepo.conn.Preload("Clubs").Find(&fixtures).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fixtures, errs
}

func (fixtureRepo *FixtureGormRepo) Fixture(id uint32) (*entity.Fixture, []error) {
	fixture := entity.Fixture{}
	errs := fixtureRepo.conn.Preload("Clubs").First(&fixture, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &fixture, errs
}

func (fixtureRepo *FixtureGormRepo) UpdateFixture(fixture *entity.Fixture) (*entity.Fixture, []error) {
	errs := fixtureRepo.conn.Save(fixture).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fixture, errs
}

func (fixtureRepo *FixtureGormRepo) DeleteFixture(id uint32) (*entity.Fixture, []error) {
	fixture, errs := fixtureRepo.Fixture(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = fixtureRepo.conn.Delete(fixture, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fixture, errs
}

func (fixtureRepo *FixtureGormRepo) StoreFixture(fixture *entity.Fixture) (*entity.Fixture, []error) {
	errs := fixtureRepo.conn.Create(fixture).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fixture, errs
}
