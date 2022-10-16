package dto

import (
	"profiler/pkg/models"
)

type ProjectReqDTO struct {
	Name        string `json:"name"`
	CompanyId   int    `json:"companyId"`
	Description string `json:"description"`
}

func (p *ProjectReqDTO) GetDTO() ProjectReqDTO {
	return *p
}

func (p *ProjectReqDTO) MapToDO(domain *models.Project) *models.Project {
	domain.Name = p.Name
	domain.Description = p.Description
	domain.CompanyID = uint(p.CompanyId)

	return domain
}

type ProjectResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Company     CompanyResponse `json:"company"`
	Deleted     bool            `json:"deleted"`
	Description string          `json:"description"`
	LastUpdated string          `json:"lastUpdated"`
}

func (p *ProjectResponse) GetDTO() ProjectResponse {
	return *p
}

func (p *ProjectResponse) MapToDO(domain *models.Project) *models.Project {
	domain.Name = p.Name
	domain.Description = p.Description
	return domain
}
