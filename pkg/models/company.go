package models

import (
	"fmt"
	"time"
)

type Company struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Founder   []Founder `json:"founders" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Project   []Project `json:"projects" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c Company) Me() string {
	companyStr := fmt.Sprintf("Company: %v founded by %v", c.Name, c.Founder)
	return companyStr
}

func (c Company) Validate() error {
	return nil
}
