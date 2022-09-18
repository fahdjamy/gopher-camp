package utils

import (
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types/dto"
)

func CompanyToCompanyRespDTO(company models.Company) dto.CompanyResponse {
	return dto.CompanyResponse{
		Name: company.Name,
	}
}
