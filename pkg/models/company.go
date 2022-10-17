package models

import (
	"errors"
	"fmt"
	"profiler/pkg/constants"
	"time"
)

type Company struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Position  string    `json:"position"`
	Founder   []Founder `json:"founders,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Project   []Project `json:"projects,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Website   string    `json:"website,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Deleted   bool      `json:"deleted,omitempty"`
}

func (c *Company) ToString() string {
	companyStr := fmt.Sprintf("Company: %v founded by %v", c.Name, c.Founder)
	return companyStr
}

func (c *Company) IsPresent() bool {
	return c.EndDate.IsZero()
}

func (c *Company) Validate() error {
	if c.Name == "" {
		return fmt.Errorf(fmt.Sprintf(constants.EmptyFieldErrorTmp, "Name"))
	}
	minSize := 2
	maxSize := 50
	if len(c.Name) < minSize || len(c.Name) > maxSize {
		return fmt.Errorf(fmt.Sprintf(constants.OutOfSizeValueErrorTmp, "Description", minSize, maxSize))
	}
	if c.Website == "" {
		return fmt.Errorf(fmt.Sprintf(constants.EmptyFieldErrorTmp, "Website"))
	}
	if len(c.Founder) == 0 {
		return errors.New("there should at least be one founder")
	}
	return nil
}

func (c *Company) Me() *Company {
	return c
}
