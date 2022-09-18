package dto

import (
	"gopher-camp/pkg/models"
)

type ProjectReqDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CompanyId   int    `json:"companyId"`
}

func (p *ProjectReqDTO) GetDTO() ProjectReqDTO {
	return *p
}

func (p *ProjectReqDTO) MapToDO(domain *models.Project) *models.Project {
	domain.Name = p.Name
	domain.Description = p.Description

	return domain
}

type ProjectResponseDTO struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Company     CompanyResponse `json:"company"`
	Description string          `json:"description"`
	LastUpdated string          `json:"lastUpdated"`
}

func (p *ProjectResponseDTO) GetDTO() ProjectResponseDTO {
	return *p
}

func (p *ProjectResponseDTO) MapToDO(domain *models.Project) *models.Project {
	domain.Name = p.Name
	domain.Description = p.Description
	return domain
}
