package services

import (
	"fmt"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage/database"
	"gopher-camp/pkg/types"
	"gorm.io/gorm"
	"strings"
	"time"
)

var coServiceErrSrc = "service.company"

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
	company, err := c.FindById(id)
	if err != nil {
		return false, err
	}

	c.db.Delete(company)
	return true, nil
}

func (c CompanyService) FindById(id uint) (*models.Company, error) {
	company := &models.Company{}
	rec := c.db.Where("id = ?", id).First(company)

	if rec.RowsAffected == 0 {
		return nil, c.createErr(nil, fmt.Sprintf("company with id (%v) does not exist", id))
	}
	return company, nil
}

func (c CompanyService) Create(model *models.Company) (*models.Company, error) {
	err := model.Validate()
	if err != nil {
		return nil, c.createErr(err, "")
	}
	if c.findByName(model.Name) != nil {
		return nil, c.createErr(err, "name must be unique")
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

func (c CompanyService) createErr(err error, msg string) error {
	return types.CustomError{
		Err:      err,
		Message:  msg,
		DateTime: time.Now(),
		Source:   fmt.Sprintf("%v.%v", coServiceErrSrc, "Create"),
	}
}

func NewCompanyService(db database.Database, logger types.Logger) CompanyService {
	return CompanyService{
		db:     db.GetDB(),
		logger: logger,
	}
}
