package models

import (
	"fmt"
	"time"
)

type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CompanyID   uint      `json:"company" gorm:"constraint:OnUpdate:CASCADE"`
}

func (p Project) ToString() string {
	me := fmt.Sprintf("%v with id: %d", p.Name, p.ID)
	return me
}

func (p Project) Validate() error {
	//TODO: implement validation logic for model
	return nil
}

func NewProject() *Project {
	return &Project{}
}
