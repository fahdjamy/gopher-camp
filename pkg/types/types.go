package types

import (
	"errors"
	"gopher-camp/pkg/models"
)

type AllServices struct {
	CompanyService Storage[models.Company]
	FounderService Storage[models.Founder]
	ProjectService Storage[models.Project]
}

func (as AllServices) Validate() error {
	if as.FounderService == nil {
		return errors.New("initialize FounderService")
	}
	if as.CompanyService == nil {
		return errors.New("initialize CompanyService")
	}
	if as.ProjectService == nil {
		return errors.New("initialize ProjectService")
	}

	return nil
}
