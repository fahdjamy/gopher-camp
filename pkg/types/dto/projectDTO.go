package dto

import (
	"gopher-camp/pkg/models"
	"time"
)

type ProjectDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CompanyId   int       `json:"companyId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (p ProjectDTO) MapToDO(domainObj *models.Project) *models.Project {
	domainObj.Name = p.Name
	domainObj.Description = p.Description

	return domainObj
}
