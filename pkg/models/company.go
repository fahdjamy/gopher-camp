package models

import "github.com/jinzhu/gorm"

type Company struct {
	gorm.Model
	ID      string `gorm:""json:"id"`
	Name    string `json:"name"`
	Founder string `json:"founder"`
	Website string `json:"website"`
}
