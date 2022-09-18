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
	coService storage.Storage[models.Company, dto.CompanyResponse, dto.CompanyResponse]
}

func (p ProjectService) FindAll() []dto.ProjectResponseDTO {
	var projects []models.Project
	p.db.Preload("Companies").Find(&projects)
	var projectsResponse []dto.ProjectResponseDTO
	for _, prj := range projects {
		projectsResponse = append(projectsResponse, p.convertToProjectResponse(prj))
	}
	return projectsResponse
}

func (p ProjectService) Delete(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectService) FindById(id int) (dto.ProjectResponseDTO, error) {
	project := &models.Project{}
	var projectResponse dto.ProjectResponseDTO
	rec := p.db.Where("id = ?", id).Limit(1).Find(project)

	if rec.RowsAffected == 0 {
		return projectResponse, errors.New(fmt.Sprintf("project with id (%v) does not exist", id))
	}

	return p.convertToProjectResponse(*project), nil
}

func (p ProjectService) Create(newProject types.DTOMapper[models.Project, dto.ProjectReqDTO]) (dto.ProjectResponseDTO, error) {
	project := models.NewProject()
	var projectResp dto.ProjectResponseDTO
	err := convertProjectDTOToProject(newProject, project)
	if err != nil {
		p.logger.LogError(err, "ProjectService.Create", "service")
		return projectResp, err
	}
	projectDto := newProject.GetDTO()
	companyResponse, err := p.coService.FindById(projectDto.CompanyId)
	if err != nil {
		return projectResp, err
	}

	project.CompanyID = companyResponse.ID

	p.db.Create(project)
	p.db.Preload("Companies").Find(project)

	projectResp.Company = companyResponse
	return projectResp, nil
}

func (p ProjectService) Update(id int, project types.DTOMapper[models.Project, dto.ProjectReqDTO]) (dto.ProjectResponseDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectService) convertToProjectResponse(project models.Project) dto.ProjectResponseDTO {
	company, _ := p.coService.FindById(int(project.CompanyID))

	return dto.ProjectResponseDTO{
		Company:     company,
		Name:        project.Name,
		ID:          project.ID,
		Description: project.Description,
	}
}

func NewProjectService(db database.Database, logger types.Logger, coService storage.Storage[models.Company, dto.CompanyResponse, dto.CompanyResponse]) ProjectService {
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
