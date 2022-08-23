package models

import (
	"github.com/jinzhu/gorm"
	"gopher-camp/pkg/config"
)

var db *gorm.DB

type Project struct {
	gorm.Model
	ID          string   `gorm:""json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Company     *Company `json:"company"`
}

func init() {
	config.OpenConnection()
	db = config.GetDB()
	db.AutoMigrate(&Project{})
}

func (p *Project) Create() *Project {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func (p *Project) GetAll() []Project {
	var projects []Project
	db.Find(&projects)
	return projects
}

func (p *Project) FindById(id int64) (*Project, *gorm.DB) {
	var project Project
	db.Where("ID=?", id).Find(&project)
	return &project, db
}

func (p *Project) DeleteById(id int64) Project {
	var project Project
	db.Where("ID=?", id).Delete(project)
	return project
}
