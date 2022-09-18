package dto

import "gopher-camp/pkg/models"

type CompanyResponse struct {
	Name        string            `json:"name"`
	Founder     []FounderResponse `json:"founder"`
	Website     string            `json:"website"`
	LastUpdated string            `json:"lastUpdated"`
}

func (c *CompanyResponse) GetDTO() CompanyResponse {
	return *c
}

func (c *CompanyResponse) MapToDO(company *models.Company) *models.Company {
	company.Name = c.Name
	company.Website = c.Website

	return company
}
