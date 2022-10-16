package dto

import "profiler/pkg/models"

type CompanyResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name,omitempty"`
	Deleted     bool              `json:"deleted"`
	Founder     []FounderResponse `json:"founder,omitempty"`
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
