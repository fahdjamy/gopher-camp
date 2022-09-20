package services

import (
	"errors"
	"fmt"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage/database"
	"gopher-camp/pkg/types"
	"gorm.io/gorm"
	"log"
	"strings"
)

type CompanyService struct {
	db     *gorm.DB
	logger types.Logger
}

func (c CompanyService) FindAll() []models.Company {
	var companies []models.Company
	c.db.Preload("Founders").Find(&companies)

	return companies
}

func (c CompanyService) Delete(id uint) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c CompanyService) FindById(id uint) (*models.Company, error) {
	log.Println(id)
	company := &models.Company{}
	rec := c.db.Where("id = ?", id).First(company)

	if rec.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("company with id (%v) does not exist", id))
	}
	return company, nil
}

func (c CompanyService) Create(model *models.Company) (*models.Company, error) {
	if c.findByName(model.Name) != nil {
		return nil, errors.New("name must be unique")
	}
	model.Name = strings.ToLower(model.Name)

	c.db.Create(model)

	return model, nil
}

func (c CompanyService) Update(id uint, model *models.Company) (*models.Company, error) {
	//TODO implement me
	panic("implement me")
}

func (c CompanyService) findByName(name string) *models.Company {
	var company models.Company
	c.db.First(&company, "name = ?", strings.ToLower(name))

	if company.ID != 0 {
		return &company
	}
	return nil
}

func NewCompanyService(db database.Database, logger types.Logger) CompanyService {
	return CompanyService{
		db:     db.GetDB(),
		logger: logger,
	}
}
