package services

import (
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage/database"
	"gopher-camp/pkg/types"
	"gorm.io/gorm"
	"time"
)

var profileServiceErrSrc = "service.company"

type ProfileService struct {
	db     *gorm.DB
	logger types.Logger
}

func (p ProfileService) FindAll() []models.Profile {
	//TODO implement me
	panic("implement me")
}

func (p ProfileService) Delete(id uint) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileService) FindById(id uint) (*models.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileService) Create(model *models.Profile) (*models.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileService) Update(id uint, model *models.Profile) (*models.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileService) createErr(err error, msg string, src string) error {
	customErr := types.CustomError{
		Err:      err,
		Message:  msg,
		Source:   src,
		DateTime: time.Now(),
	}
	if src == "" {
		customErr.Source = profileServiceErrSrc
	}
	return customErr
}

func NewProfileService(db database.Database, logger types.Logger) ProfileService {
	return ProfileService{
		db:     db.GetDB(),
		logger: logger,
	}
}
