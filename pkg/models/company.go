package models

import (
	"errors"
	"fmt"
	"gopher-camp/pkg/constants"
	"time"
)

type Company struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Founder   []Founder `json:"founders,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Project   []Project `json:"projects,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Website   string    `json:"website,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (c *Company) ToString() string {
	companyStr := fmt.Sprintf("Company: %v founded by %v", c.Name, c.Founder)
	return companyStr
}

func (c *Company) Validate() error {
	if c.Name == "" {
		return errors.New(fmt.Sprintf(constants.EmptyFieldErrorTmp, "Name"))
	}
	minSize := 2
	maxSize := 50
	if len(c.Name) < minSize || len(c.Name) > maxSize {
		return errors.New(fmt.Sprintf(constants.OutOfSizeValueErrorTmp, "Description", minSize, maxSize))
	}
	if c.Website == "" {
		return errors.New(fmt.Sprintf(constants.EmptyFieldErrorTmp, "Website"))
	}
	if len(c.Founder) == 0 {
		return errors.New("there should at least be one founder")
	}
	return nil
}

func (c *Company) Me() *Company {
	return c
}
