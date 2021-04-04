package fixture

import "github.com/solkn/soccer/api/entity"

type FixturesRepository interface {
	Fixtures() ([]entity.Fixture, []error)
	Fixture(id uint32) (*entity.Fixture, []error)
	StoreFixture(user *entity.Fixture) (*entity.Fixture, []error)
	UpdateFixture(order *entity.Fixture) (*entity.Fixture, []error)
	DeleteFixture(id uint32) (*entity.Fixture, []error)
}
