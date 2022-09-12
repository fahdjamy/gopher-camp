package models

import (
	"fmt"
	"time"
)

type Company struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Founder   Founder   `json:"founder"`
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c Company) Me() string {
	companyStr := fmt.Sprintf("Company: %v founded by %v", c.Name, c.Founder)
	return companyStr
}

func (c Company) Validate() error {
	//TODO implement me
	panic("implement me")
}
