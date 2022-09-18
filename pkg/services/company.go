package services

import (
	"errors"
	"fmt"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/dto"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types"
	"gorm.io/gorm"
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

func (c CompanyService) Delete(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c CompanyService) FindById(id int) (*models.Company, error) {
	company := &models.Company{}
	rec := c.db.Where("id = ?", id).First(company)

	if rec.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("company with id (%v) does not exist", id))
	}
	return company, nil
}

func (c CompanyService) Create(model types.DTOMapper[models.Company, dto.ProjectReqDTO]) (*models.Company, error) {
	//TODO implement me
	panic("implement me")
}

func (c CompanyService) Update(id int, model types.DTOMapper[models.Company, dto.ProjectReqDTO]) (*models.Company, error) {
	//TODO implement me
	panic("implement me")
}

func NewCompanyService(db database.Database, logger types.Logger) CompanyService {
	return CompanyService{
		db:     db.GetDB(),
		logger: logger,
	}
}
