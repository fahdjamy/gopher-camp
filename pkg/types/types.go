package types

import (
	"errors"
	"fmt"
	"gopher-camp/pkg/models"
	"time"
)

type AllServices struct {
	CompanyService DOServiceProvider[models.Company]
	FounderService DOServiceProvider[models.Founder]
	ProjectService DOServiceProvider[models.Project]
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

type CustomError struct {
	Err      error
	Source   string
	Message  string
	DateTime time.Time
}

func (e CustomError) Error() string {
	errStr := ""
	if e.Err != nil {
		errStr = errStr + e.Err.Error()
	}
	return fmt.Sprintf("[source]:%v [error]: %v", e.Source, errStr)
}

func NewCustomError() *CustomError {
	return &CustomError{
		Err:      nil,
		Source:   "hidden",
		Message:  "error",
		DateTime: time.Now(),
	}
}
