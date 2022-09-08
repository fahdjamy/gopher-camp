package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types"
)

type ProjectService struct {
	db *gorm.DB
}

func (p ProjectService) FindAll() []types.Domain[models.Project] {
	var projects []models.Project
	p.db.Find(&projects)

	var domainProjects []types.Domain[models.Project]

	fmt.Println(len(projects))

	for _, prj := range projects {
		domainProjects = append(domainProjects, models.ProjectToDomainProject(&prj))
	}

	return domainProjects
}

func (p ProjectService) Delete(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectService) FindById(id int) (*types.Domain[models.Project], error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectService) Create(model types.Domain[models.Project]) (*types.Domain[models.Project], error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectService) Update(id int, model types.Domain[models.Project]) (*types.Domain[models.Project], error) {
	//TODO implement me
	panic("implement me")
}

func NewProjectService(db database.Database) ProjectService {
	return ProjectService{
		db: db.GetDB(),
	}
}
