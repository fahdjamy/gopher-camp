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

func (c CompanyService) FindAll() []dto.CompanyResponse {
	var companies []models.Company
	c.db.Preload("Founders").Find(&companies)

	var companiesResponse []dto.CompanyResponse

	for _, co := range companies {
		companiesResponse = append(companiesResponse, c.convertToCompanyResponse(co))
	}

	return companiesResponse
}

func (c CompanyService) Delete(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c CompanyService) FindById(id int) (dto.CompanyResponse, error) {
	company := &models.Company{}
	rec := c.db.Where("id = ?", id).First(company)

	if rec.RowsAffected == 0 {
		return dto.CompanyResponse{}, errors.New(fmt.Sprintf("company with id (%v) does not exist", id))
	}
	return c.convertToCompanyResponse(*company), nil
}

func (c CompanyService) Create(model types.DTOMapper[models.Company, dto.CompanyResponse]) (dto.CompanyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CompanyService) Update(id int, model types.DTOMapper[models.Company, dto.CompanyResponse]) (dto.CompanyResponse, error) {
	//TODO implement me
	panic("implement me")
}
func (c CompanyService) convertToCompanyResponse(company models.Company) dto.CompanyResponse {
	return dto.CompanyResponse{
		ID:      company.ID,
		Name:    company.Name,
		Website: company.Website,
	}
}

func NewCompanyService(db database.Database, logger types.Logger) CompanyService {
	return CompanyService{
		db:     db.GetDB(),
		logger: logger,
	}
}
