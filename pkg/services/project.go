package services

import (
	"errors"
	"fmt"
	"gopher-camp/pkg/config/database"
	"gopher-camp/pkg/dto"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage"
	"gopher-camp/pkg/types"
	"gorm.io/gorm"
)

type ProjectService struct {
	db        *gorm.DB
	logger    types.Logger
	coService storage.Storage[models.Company, dto.ProjectReqDTO]
}

func (p ProjectService) FindAll() []models.Project {
	var projects []models.Project
	p.db.Preload("Companies").Find(&projects)
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

func (p ProjectService) Create(newProject types.DTOMapper[models.Project, dto.ProjectReqDTO]) (*models.Project, error) {
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
	project.CompanyID = company.ID
	p.db.Create(project)
	p.db.Preload("Companies").Find(project)

	return project, nil
}

func (p ProjectService) Update(id int, project types.DTOMapper[models.Project, dto.ProjectReqDTO]) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func NewProjectService(db database.Database, logger types.Logger, coService storage.Storage[models.Company, dto.ProjectReqDTO]) ProjectService {
	return ProjectService{
		db:        db.GetDB(),
		logger:    logger,
		coService: coService,
	}
}

func convertProjectDTOToProject(projectDTO types.DTOMapper[models.Project, dto.ProjectReqDTO], project *models.Project) error {
	projectDTO.MapToDO(project)

	return nil
}
