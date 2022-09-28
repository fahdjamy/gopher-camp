package models

import (
	"errors"
	"fmt"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/validators"
	"time"
)

type Founder struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Deleted   bool      `json:"deleted"`
	CompanyID uint      `gorm:"constraint:OnUpdate:CASCADE"`
	Email     string    `json:"founder"`
	LinkedIn  string    `json:"linkedIn"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (f *Founder) ToString() string {
	str := fmt.Sprintf("Name: %v LinkedIn by %v", f.Name, f.LinkedIn)
	return str
}

func (f *Founder) Validate() error {
	if f.Name == "" {
		return errors.New(fmt.Sprintf(constants.EmptyFieldErrorTmp, "Name"))
	}
	minSize := 2
	maxSize := 150
	if len(f.Name) < minSize || len(f.Name) > maxSize {
		return errors.New(fmt.Sprintf(fmt.Sprintf(constants.OutOfSizeValueErrorTmp, "Name", minSize, maxSize)))
	}
	if f.Email != "" {
		_, err := validators.ValidateEmail(f.Email)
		if err != nil {
			return errors.New("invalid Email")
		}
	}
	if f.LinkedIn == "" {
		return errors.New(fmt.Sprintf(constants.EmptyFieldErrorTmp, "LinkedIn"))
	}
	return nil
}

func (f *Founder) Me() *Founder {
	return f
}

func NewFounder() *Founder {
	return &Founder{}
}
