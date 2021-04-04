package services

import (
	"github.com/solkn/soccer/api/entity"
	fixture "github.com/solkn/soccer/api/fixture"
)

type FixtureService struct {
	fixtureRepo fixture.FixturesRepository
}

func NewClubService(clubRepo fixture.FixturesRepository) *FixtureService {
	return &FixtureService{fixtureRepo: clubRepo}
}

func (fs *FixtureService) Fixtures() ([]entity.Fixture, []error) {

	fixtures, errs := fs.fixtureRepo.Fixtures()
	if len(errs) > 0 {
		return nil, errs
	}
	return fixtures, errs

}

func (fs *FixtureService) Fixture(id uint32) (*entity.Fixture, []error) {
	fxtr, errs := fs.fixtureRepo.Fixture(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return fxtr, errs

}

func (fs *FixtureService) UpdateFixture(fxt *entity.Fixture) (*entity.Fixture, []error) {
	fxtr, errs := fs.fixtureRepo.UpdateFixture(fxt)
	if len(errs) > 0 {
		return nil, errs
	}
	return fxtr, errs

}

func (fs *FixtureService) DeleteFixture(id uint32) (*entity.Fixture, []error) {

	fxtr, errs := fs.fixtureRepo.DeleteFixture(uint32(id))
	if len(errs) > 0 {
		return nil, errs
	}
	return fxtr, errs
}

func (fs *FixtureService) StoreFixture(fxt *entity.Fixture) (*entity.Fixture, []error) {

	fxtr, errs := fs.fixtureRepo.StoreFixture(fxt)
	if len(errs) > 0 {
		return nil, errs
	}
	return fxtr, errs
}
