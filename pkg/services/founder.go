package services

import (
	"errors"
	"fmt"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage/database"
	"gopher-camp/pkg/types"
	"gorm.io/gorm"
	"strings"
)

type FounderService struct {
	db     *gorm.DB
	logger types.Logger
}

func (f FounderService) FindAll() []models.Founder {
	var founders []models.Founder
	f.db.Find(&founders)

	return founders
}

func (f FounderService) Delete(id uint) (bool, error) {
	founder, err := f.FindById(id)
	if err != nil {
		return false, err
	}
	f.db.Delete(founder)

	return true, nil
}

func (f FounderService) FindById(id uint) (*models.Founder, error) {
	founder := &models.Founder{}
	rec := f.db.Where("id = ?", id).First(founder)

	if rec.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("company with id (%v) does not exist", id))
	}
	return founder, nil
}

func (f FounderService) Create(founder *models.Founder) (*models.Founder, error) {
	if f.findByName(founder.Name) != nil {
		return nil, errors.New(fmt.Sprintf("Founder name must be unique, Duplicate name %v", founder.Name))
	}
	if f.findByEmail(founder.Email) != nil {
		return nil, errors.New(fmt.Sprintf("Founder Email must be unique, Duplicate Email %v", founder.Email))
	}
	founder.Name = strings.ToLower(founder.Name)
	f.db.Create(founder)

	return founder, nil
}

func (f FounderService) findByName(name string) *models.Founder {
	var founder models.Founder
	f.db.First(&founder, "name = ?", strings.ToLower(name))

	if founder.Name != "" {
		return &founder
	}
	return nil
}

func (f FounderService) findByEmail(email string) *models.Founder {
	var founder models.Founder
	f.db.First(&founder, "email = ?", email)

	if founder.ID != 0 {
		return &founder
	}
	return nil
}

func (f FounderService) Update(id uint, model *models.Founder) (*models.Founder, error) {
	founder, err := f.FindById(id)
	if err != nil {
		return nil, err
	}
	founder.Email = model.Email
	return founder, nil
}

func NewFounderService(db database.Database, logger types.Logger) FounderService {
	return FounderService{
		db:     db.GetDB(),
		logger: logger,
	}
}
