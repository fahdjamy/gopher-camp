package helpers

import (
	"encoding/json"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types"
	"gopher-camp/pkg/types/dto"
	"gopher-camp/pkg/utils"
)

func SeedDatabaseData(logger types.Logger, services types.AllServices) error {
	if err := services.Validate(); err != nil {
		return err
	}
	foundersData, err := readFounders()
	if err != nil {
		logger.LogError(err, "utils.ReadJsonFile", "utils")
		return err
	}
	if foundersData != nil && len(foundersData) > 0 {
		for _, fd := range foundersData {
			_, err := services.FounderService.Create(fd)
			if err != nil {
				return err
			}
		}
	}

	companiesData, err := readCompanies(foundersData)
	if err != nil {
		logger.LogError(err, "utils.ReadJsonFile", "utils")
		return err
	}

	if companiesData != nil && len(companiesData) > 0 {
		for _, company := range companiesData {
			_, err := services.CompanyService.Create(company)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func readFounders() ([]*models.Founder, error) {
	var data []dto.FounderRequest
	jsonFounders, err := utils.ReadJsonFile(utils.AbsPathToProject("./local/data/founders.json"))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonFounders, &data)

	var founders []*models.Founder
	for _, founder := range data {
		founders = append(founders, &models.Founder{
			Name:     founder.Name,
			Email:    founder.Email,
			LinkedIn: founder.LinkedIn,
		})
	}

	return founders, err
}

func readCompanies(founders []*models.Founder) ([]*models.Company, error) {
	var data []dto.CompanyRequest
	jsonCompanies, err := utils.ReadJsonFile(utils.AbsPathToProject("./local/data/companies.json"))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonCompanies, &data)

	var companies []*models.Company

	for _, company := range data {
		var companyFounders []models.Founder
		for _, founderIndex := range company.Founders {
			companyFounders = append(companyFounders, *findFounder(founderIndex, founders))
		}
		companies = append(companies, &models.Company{
			Name:    company.Name,
			Website: company.Website,
			Founder: companyFounders,
		})
	}

	return companies, err
}

func findFounder(index int, founders []*models.Founder) *models.Founder {
	for idx, fd := range founders {
		if idx+1 == index {
			return fd
		}
	}

	return nil
}
