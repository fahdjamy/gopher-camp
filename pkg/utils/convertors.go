package utils

import (
	"profiler/pkg/models"
	"profiler/pkg/types/dto"
)

func CompanyToCompanyRespDTO(company models.Company) dto.CompanyResponse {
	return dto.CompanyResponse{
		Name: company.Name,
	}
}
