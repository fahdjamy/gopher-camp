package models

import (
	"fmt"
	"gopher-camp/pkg/validators"
	"time"
)

type Founder struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"founder"`
	LinkedIn  string    `json:"website"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (f Founder) Me() string {
	str := fmt.Sprintf("Name: %v LinkedIn by %v", f.Name, f.LinkedIn)
	return str
}

func (f Founder) Validate() error {
	_, err := validators.ValidateEmail(f.Email)
	if err != nil {
		return err
	}
	return nil
}
