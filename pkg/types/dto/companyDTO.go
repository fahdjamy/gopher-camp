package dto

import "gopher-camp/pkg/models"

type CompanyResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Founder     []FounderResponse `json:"founder"`
	Website     string            `json:"website"`
	LastUpdated string            `json:"lastUpdated"`
}

type CompanyRequest struct {
	Name     string `json:"name"`
	Founders []int  `json:"founders"`
	Website  string `json:"website"`
}

func (c *CompanyResponse) GetDTO() CompanyResponse {
	return *c
}

func (c *CompanyResponse) MapToDO(company *models.Company) *models.Company {
	company.Name = c.Name
	company.Website = c.Website

	return company
}
