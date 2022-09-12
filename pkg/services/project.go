package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage"
	"gopher-camp/pkg/types/dto"
	"gopher-camp/pkg/utils"
)

type ProjectService struct {
	db        *gorm.DB
	logger    utils.CustomLogger
	coService storage.Storage[models.Company]
}

func (p ProjectService) FindAll() []models.Project {
	var projects []models.Project
	p.db.Find(&projects)
	return projects
}

func (p ProjectService) Delete(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectService) FindById(id int) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectService) Create(newProject dto.ToDomainMapper[models.Project]) (*models.Project, error) {
	project := models.NewProject()
	err := convertProjectDTOToProject(newProject, project)
	if err != nil {
		p.logger.LogError(err, "ProjectService.Create", "service")
		return nil, err
	}
	projectDto := newProject.(dto.ProjectDTO)
	company, err := p.coService.FindById(projectDto.CompanyId)
	if company == nil {
		return nil, fmt.Errorf("company with ID %v does not exist", projectDto.CompanyId)
	}
	project.Company = company
	p.db.NewRecord(project)

	return project, nil
}

func (p ProjectService) Update(id int, project dto.ToDomainMapper[models.Project]) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func NewProjectService(db database.Database) ProjectService {
	return ProjectService{
		db: db.GetDB(),
	}
}

func convertProjectDTOToProject(projectDTO dto.ToDomainMapper[models.Project], project *models.Project) error {
	projectDTO.MapToDO(project)

	return nil
}
