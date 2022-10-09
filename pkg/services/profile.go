package services

import (
	"fmt"
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
	profile, err := p.FindById(id)
	if err != nil {
		return false, err
	}

	profile.Deleted = true
	_, err = p.Update(id, profile)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p ProfileService) FindById(id uint) (*models.Profile, error) {
	profile := &models.Profile{}

	rec := p.db.Where("id = ?", id).Limit(1).Find(profile)

	if rec.RowsAffected == 0 {
		return nil, p.createErr(nil,
			fmt.Sprintf("profile with id (%v) does not exist", id),
			fmt.Sprintf("%v.FindById", profileServiceErrSrc),
		)
	}
	return profile, nil
}

func (p ProfileService) Create(profile *models.Profile) (*models.Profile, error) {
	err := profile.Validate()
	if err != nil {
		return nil, err
	}
	if storedProfile, _ := p.FindByEmail(profile.Email); storedProfile != nil {
		return nil, p.createErr(nil, "email is taken", fmt.Sprintf("%v.Create", profileServiceErrSrc))
	}
	if storedProfile, _ := p.FindByNickname(profile.Nickname); storedProfile != nil {
		return nil, p.createErr(nil, "nickname is taken", fmt.Sprintf("%v.Create", profileServiceErrSrc))
	}

	p.db.Create(profile)
	p.db.Find(profile)

	return profile, nil
}

func (p ProfileService) Update(id uint, profile *models.Profile) (*models.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileService) FindByNickname(nickname string) (*models.Profile, error) {
	profile := &models.Profile{}

	rec := p.db.Where("nickname = ?", nickname).Limit(1).Find(profile)

	if rec.RowsAffected == 0 {
		return nil, p.createErr(nil, fmt.Sprintf("profile with nickname (%v) does not exist", nickname), "")
	}

	return profile, nil
}

func (p ProfileService) FindByEmail(email string) (*models.Profile, error) {
	profile := &models.Profile{}

	rec := p.db.Where("email = ?", email).Limit(1).Find(profile)

	if rec.RowsAffected == 0 {
		return nil, p.createErr(nil, fmt.Sprintf("email is taken (%v) exist", email), "")
	}

	return profile, nil
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
