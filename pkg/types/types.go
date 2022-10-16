package types

import (
	"errors"
	"fmt"
	"profiler/pkg/models"
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
	Err      error     `json:"error,omitempty"`
	Source   string    `json:"source,omitempty"`
	Message  string    `json:"errorMessage,omitempty"`
	DateTime time.Time `json:"dateTime,omitempty"`
}

func (e CustomError) Error() string {
	errStr := e.Message
	if errStr == "" && e.Err != nil {
		errStr = e.Err.Error()
	}
	return fmt.Sprintf(errStr)
}

func NewCustomError() *CustomError {
	return &CustomError{
		Err:      nil,
		Source:   "hidden",
		Message:  "error",
		DateTime: time.Now(),
	}
}
