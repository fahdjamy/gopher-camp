package utils

import (
	"gopher-camp/pkg/dto"
	"gopher-camp/pkg/models"
)

func CompanyToCompanyRespDTO(company models.Company) dto.CompanyResponse {
	return dto.CompanyResponse{
		Name: company.Name,
	}
}
