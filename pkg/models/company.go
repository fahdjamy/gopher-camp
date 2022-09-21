package models

import (
	"fmt"
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

func (c Company) ToString() string {
	companyStr := fmt.Sprintf("Company: %v founded by %v", c.Name, c.Founder)
	return companyStr
}

func (c Company) Validate() error {
	return nil
}
