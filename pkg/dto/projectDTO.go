package dto

import (
	"gopher-camp/pkg/models"
)

type ProjectReqDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CompanyId   int    `json:"companyId"`
}

type ProjectResponseDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LastUpdated string `json:"lastUpdated"`
}

func (p *ProjectReqDTO) GetDTO() ProjectReqDTO {
	return *p
}

func (p *ProjectReqDTO) MapToDO(domain *models.Project) *models.Project {
	domain.Name = p.Name
	domain.Description = p.Description

	return domain
}
