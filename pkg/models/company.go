package models

import (
	"time"
)

type Company struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Founder   string    `json:"founder"`
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
