package models

import (
	"errors"
	"fmt"
	"profiler/pkg/constants"
	"profiler/pkg/validators"
	"time"
)

type Profile struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Bio       string    `json:"bio,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	LinkedIn  string    `json:"linkedIn,omitempty"`
	Facebook  string    `json:"facebook,omitempty"`
	Twitter   string    `json:"twitter,omitempty"`
	Youtube   string    `json:"youtube',omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Deleted   bool      `json:"deleted"`
}

func (prof *Profile) ToString() string {
	profileStr := fmt.Sprintf("Fullname: %v %v", prof.FirstName, prof.LastName)
	return profileStr
}

func (prof *Profile) Validate() error {
	if prof.FirstName == "" {
		return errors.New(fmt.Sprintf(constants.EmptyFieldErrorTmp, "FirstName"))
	}
	if prof.LastName == "" {
		return errors.New(fmt.Sprintf(constants.EmptyFieldErrorTmp, "LastName"))
	}
	minSize := 2
	maxSize := 50
	if len(prof.FirstName) < minSize || len(prof.FirstName) > maxSize {
		return errors.New(fmt.Sprintf(constants.OutOfSizeValueErrorTmp, "firstName", minSize, maxSize))
	}
	if len(prof.LastName) < minSize || len(prof.LastName) > maxSize {
		return errors.New(fmt.Sprintf(constants.OutOfSizeValueErrorTmp, "lastName", minSize, maxSize))
	}
	if prof.Email == "" {
		return errors.New(fmt.Sprintf(constants.EmptyFieldErrorTmp, "Email"))
	}
	_, err := validators.ValidateEmail(prof.Email)
	if err != nil {
		return errors.New("invalid Email")
	}

	minBioSize := 50
	maxBioSize := 1000
	if prof.Bio == "" {
		return errors.New(fmt.Sprintf(constants.EmptyFieldErrorTmp, "Bio"))
	}
	if len(prof.Bio) < minBioSize || len(prof.Bio) > maxBioSize {
		return errors.New(fmt.Sprintf(constants.OutOfSizeValueErrorTmp, "bio", minSize, maxSize))
	}

	return nil
}

func (prof *Profile) Me() *Profile {
	return prof
}
