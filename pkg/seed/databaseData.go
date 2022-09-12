package seed

import (
	"errors"
	"github.com/jinzhu/gorm"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/validators"
)

type DatabaseSeeder struct {
	db *gorm.DB
}

func newCompany(founder *models.Founder) *models.Company {
	company := &models.Company{
		Founder: *founder,
		Name:    "Hacker Bay",
		Website: "https://hackerbay.io/",
	}
	return company
}

func newFounder() *models.Founder {
	founder := &models.Founder{
		Name:     "Nawaz Dhandal",
		Email:    "nawazdhandala@outlook.com",
		LinkedIn: "https://www.linkedin.com/in/nawazdhandala/",
	}
	return founder
}

func (ds DatabaseSeeder) CreateCompany(founder *models.Founder) (*models.Company, error) {
	company := newCompany(founder)
	if ds.companyExists(company) {
		return company, nil
	}
	if !ds.founderExists(founder) {
		return nil, errors.New("founder does not exist in the Database")
	}
	ds.db.Create(company)

	if validators.IsIntEmpty(company.ID) {
		return nil, errors.New("can't create company")
	}

	return company, nil
}

func (ds DatabaseSeeder) CreateFounder() (*models.Founder, error) {
	founder := newFounder()
	if ds.founderExists(founder) {
		return founder, errors.New("founder exists")
	}
	ds.db.Create(founder)
	if validators.IsIntEmpty(founder.ID) {
		return nil, errors.New("can't create founder")
	}

	return founder, nil
}

func (ds DatabaseSeeder) founderExists(founder *models.Founder) bool {
	rec := ds.db.Where("email = ?", founder.Email).Limit(1).Find(founder)

	if rec.RowsAffected > 0 {
		return true
	}
	return false
}

func (ds DatabaseSeeder) companyExists(company *models.Company) bool {
	rec := ds.db.Where("name = ?", company.Name).Limit(1).Find(company)

	if rec.RowsAffected > 0 {
		return true
	}
	return false
}

func NewDatabaseSeeder(db *gorm.DB) *DatabaseSeeder {
	return &DatabaseSeeder{
		db: db,
	}
}
