package dto

import (
	"gopher-camp/pkg/models"
)

type ProjectDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CompanyId   int    `json:"companyId"`
}

func (p *ProjectDTO) GetDTO() ProjectDTO {
	return *p
}

func (p *ProjectDTO) MapToDO(domain *models.Project) *models.Project {
	domain.Name = p.Name
	domain.Description = p.Description

	return domain
}
