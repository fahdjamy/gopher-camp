package services

import (
	"fmt"
	"gorm.io/gorm"
	"profiler/pkg/models"
	"profiler/pkg/storage/database"
	"profiler/pkg/types"
	"strings"
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
		return false, err
	}
	if project.Deleted {
		return true, nil
	}

	project.Deleted = true
	_, _ = p.Update(id, project)

	return true, nil
}

func (p ProjectService) FindById(id uint) (*models.Project, error) {
	project := &models.Project{}
	rec := p.db.Where("id = ?", id).Limit(1).Find(project)

	if rec.RowsAffected == 0 {
		return nil, p.createErr(
			nil,
			fmt.Sprintf("project with id (%v) does not exist", id),
			fmt.Sprintf("%v.%v", prjSrvErrSrc, "FindById"),
		)
	}

	return project, nil
}

func (p ProjectService) FindByName(name string) (*models.Project, error) {
	if name == "" {
		return nil, p.createErr(
			nil,
			"name is empty",
			fmt.Sprintf("%v.%v", prjSrvErrSrc, "FindByName"),
		)
	}

	project := &models.Project{}
	rec := p.db.Where("name = ?", strings.ToLower(name)).Limit(1).Find(project)
	if rec.RowsAffected == 0 {
		return nil, p.createErr(
			nil,
			fmt.Sprintf("project with name (%v) does not exist", name),
			fmt.Sprintf("%v.%v", prjSrvErrSrc, "FindByName"),
		)
	}

	return project, nil
}

func (p ProjectService) Create(project *models.Project) (*models.Project, error) {
	err := project.Validate()
	if err != nil {
		return nil, p.createErr(err, "", "models.project.Validate")
	}
	company, err := p.coService.FindById(project.CompanyID)
	if err != nil {
		return nil, p.createErr(err, "", fmt.Sprintf("%v.%v", prjSrvErrSrc, "Create"))
	}

	existingProject, _ := p.FindByName(project.Name)
	if existingProject != nil {
		return nil, p.createErr(
			nil,
			fmt.Sprintf("name must be unique. project with name (%v) already exists", project.Name),
			fmt.Sprintf("%v.%v", prjSrvErrSrc, "Create"),
		)
	}

	project.CompanyID = company.ID
	project.Name = strings.ToLower(project.Name)

	p.db.Create(project)
	p.db.Find(project)

	return project, nil
}

func (p ProjectService) Update(id uint, project *models.Project) (*models.Project, error) {
	err := project.Validate()
	if err != nil {
		return nil, p.createErr(err, "", "models.project.Validate")
	}
	storedProject, err := p.FindById(id)
	if err != nil {
		return nil, err
	}

	if storedProject.Name != project.Name {
		existingProject, _ := p.FindByName(project.Name)
		if !project.Deleted && existingProject != nil && !existingProject.Deleted {
			return nil, p.createErr(
				nil,
				fmt.Sprintf("name must be unique. project with name (%v) already exists", project.Name),
				fmt.Sprintf("%v.%v", prjSrvErrSrc, "Update"),
			)
		}
	}

	updateColumnData := map[string]interface{}{
		"name":        project.Name,
		"deleted":     project.Deleted,
		"description": project.Description,
	}

	if project.CompanyID != 0 {
		company, err := p.coService.FindById(project.CompanyID)
		if err != nil {
			return nil, err
		}
		updateColumnData["company_id"] = company.ID
	}

	p.db.Model(storedProject).Where("deleted = ?", false).Updates(updateColumnData)

	return storedProject, nil
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
