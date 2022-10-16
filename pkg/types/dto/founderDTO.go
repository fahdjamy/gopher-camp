package dto

import "profiler/pkg/models"

type FounderResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	LinkedIn    string `json:"linkedIn"`
	LastUpdated string `json:"lastUpdated"`
}

type FounderRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email'"`
	LinkedIn string `json:"linkedIn"`
}

func (f *FounderRequest) GetDTO() FounderRequest {
	return *f
}

func (f *FounderRequest) MapToDO(domain *models.Founder) *models.Founder {
	domain.Name = f.Name
	domain.Email = f.Email
	domain.LinkedIn = f.LinkedIn

	return domain
}
