package models

import (
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/types"
	"time"
)

type Project struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Company     *Company  `json:"company"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (p *Project) Create(db database.Database) *Project {
	db.GetDB().NewRecord(p)
	db.GetDB().Create(&p)
	return p
}

func (p *Project) FindAll(db database.Database) []Project {
	var projects []Project
	db.GetDB().Find(&projects)
	return projects
}

func (p *Project) FindById(id int64, db database.Database) *Project {
	var project Project
	db.GetDB().Where("ID=?", id).Find(&project)
	return &project
}

func (p *Project) DeleteById(id int64, db database.Database) Project {
	var project Project
	db.GetDB().Where("ID=?", id).Delete(project)
	return project
}

func NewProject() types.Domain[Project] {
	return &Project{}
}

func ProjectToDomainProject(project *Project) types.Domain[Project] {
	return project
}
