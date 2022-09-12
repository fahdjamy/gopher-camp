package services

import (
	"github.com/jinzhu/gorm"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types"
	"gopher-camp/pkg/types/dto"
)

type CompanyService struct {
	db     *gorm.DB
	logger types.Logger
}

func (c CompanyService) FindAll() []models.Company {
	//TODO implement me
	panic("implement me")
}

func (c CompanyService) Delete(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c CompanyService) FindById(id int) (*models.Company, error) {

	//TODO implement me
	panic("implement me")
}

func (c CompanyService) Create(model dto.ToDomainMapper[models.Company]) (*models.Company, error) {
	//TODO implement me
	panic("implement me")
}

func (c CompanyService) Update(id int, model dto.ToDomainMapper[models.Company]) (*models.Company, error) {
	//TODO implement me
	panic("implement me")
}

func NewCompanyService(db database.Database, logger types.Logger) CompanyService {
	return CompanyService{
		db:     db.GetDB(),
		logger: logger,
	}
}
