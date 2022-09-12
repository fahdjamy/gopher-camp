package services

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/dto"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage"
	"gopher-camp/pkg/types"
)

type ProjectService struct {
	db        *gorm.DB
	logger    types.Logger
	coService storage.Storage[models.Company, dto.ProjectDTO]
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
	project := &models.Project{}
	rec := p.db.Where("id = ?", id).Limit(1).Find(project)

	if rec.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("project with id (%v) does not exist", id))
	}
	return project, nil
}

func (p ProjectService) Create(newProject types.DTOMapper[models.Project, dto.ProjectDTO]) (*models.Project, error) {
	project := models.NewProject()
	err := convertProjectDTOToProject(newProject, project)
	if err != nil {
		p.logger.LogError(err, "ProjectService.Create", "service")
		return nil, err
	}
	projectDto := newProject.GetDTO()
	company, err := p.coService.FindById(projectDto.CompanyId)
	if err != nil {
		return nil, err
	}
	project.Company = company
	p.db.Create(project)

	return project, nil
}

func (p ProjectService) Update(id int, project types.DTOMapper[models.Project, dto.ProjectDTO]) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func NewProjectService(db database.Database, logger types.Logger, coService storage.Storage[models.Company, dto.ProjectDTO]) ProjectService {
	return ProjectService{
		db:        db.GetDB(),
		logger:    logger,
		coService: coService,
	}
}

func convertProjectDTOToProject(projectDTO types.DTOMapper[models.Project, dto.ProjectDTO], project *models.Project) error {
	projectDTO.MapToDO(project)

	return nil
}
