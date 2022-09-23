package services

import (
	"fmt"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage/database"
	"gopher-camp/pkg/types"
	"gopher-camp/pkg/types/dto"
	"gorm.io/gorm"
	"time"
)

var prjSrvErrSrc = "service.projectServices"

type ProjectService struct {
	db        *gorm.DB
	logger    types.Logger
	coService types.DOServiceProvider[models.Company]
}

func (p ProjectService) FindAll() []models.Project {
	var projects []models.Project
	p.db.Find(&projects)

	return projects
}

func (p ProjectService) Delete(id uint) (bool, error) {
	project, err := p.FindById(id)
	if err != nil {
		return false, p.createErr(err, "", "")
	}
	p.db.Delete(project)

	return true, nil
}

func (p ProjectService) FindById(id uint) (*models.Project, error) {
	project := &models.Project{}
	rec := p.db.Where("id = ?", id).Limit(1).Find(project)

	if rec.RowsAffected == 0 {
		return nil, p.createErr(
			nil,
			fmt.Sprintf("project with id (%v) does not exist", id),
			fmt.Sprintf("%v.%v", prjSrvErrSrc, "Create"),
		)
	}

	return project, nil
}

func (p ProjectService) Create(project *models.Project) (*models.Project, error) {
	err := project.Validate()
	if err != nil {
		return nil, p.createErr(err, "", fmt.Sprintf("%v.%v", prjSrvErrSrc, "Create"))
	}
	company, err := p.coService.FindById(project.CompanyID)
	if err != nil {
		return nil, p.createErr(err, "", fmt.Sprintf("%v.%v", prjSrvErrSrc, "Create"))
	}

	project.CompanyID = company.ID

	p.db.Create(project)
	p.db.Find(project)

	return project, nil
}

func (p ProjectService) Update(id uint, project *models.Project) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectService) createErr(err error, msg string, src string) error {
	customErr := types.CustomError{
		Err:      err,
		Source:   src,
		Message:  msg,
		DateTime: time.Now(),
	}
	if src == "" {
		customErr.Source = prjSrvErrSrc
	}
	return customErr
}

func NewProjectService(db database.Database, logger types.Logger, coService types.DOServiceProvider[models.Company]) ProjectService {
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
