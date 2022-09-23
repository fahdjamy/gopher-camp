package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types"
	"gopher-camp/pkg/types/dto"
	"gopher-camp/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

type ProjectController struct {
	service  types.DOServiceProvider[models.Project]
	services types.AllServices
}

func (pc ProjectController) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects := pc.service.FindAll()
	var projectsResponse []dto.ProjectResponse
	for _, project := range projects {
		projectsResponse = append(projectsResponse, pc.convertToProject(project))
	}

	res, _ := json.Marshal(utils.SuccessArray(projectsResponse, "success"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (pc ProjectController) DeleteProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectId := params["projectId"]
	_, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//project := dbProject.DeleteById(id, pc.db)

	w.WriteHeader(http.StatusOK)
	//_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}

func (pc ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (pc ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	projectDO := models.NewProject()
	projectDTO := &dto.ProjectReqDTO{}
	utils.ParseBody(r, projectDTO)
	log.Println(projectDTO.CompanyId)

	project, err := pc.service.Create(projectDTO.MapToDO(projectDO))
	if err != nil {
		res, _ := json.Marshal(utils.CreateFailureWithMessage(err.(types.CustomError)))
		_, _ = w.Write(res)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(utils.SingleObject(pc.convertToProject(*project)))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (pc ProjectController) GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectId := params["projectId"]
	id, err := strconv.ParseUint(projectId, 0, 0)
	if err != nil {
		res, _ := json.Marshal(utils.CreateFailure(types.CustomError{Err: err}))
		_, _ = w.Write(res)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	project, err := pc.service.FindById(uint(id))
	if err != nil {
		res, _ := json.Marshal(utils.CreateFailureWithMessage(err.(types.CustomError)))
		_, _ = w.Write(res)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(utils.SingleObject(pc.convertToProject(*project)))
	_, _ = w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func (pc ProjectController) convertToProject(project models.Project) dto.ProjectResponse {
	company, _ := pc.services.CompanyService.FindById(project.CompanyID)

	return dto.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		Company: dto.CompanyResponse{
			ID:          company.ID,
			Name:        company.Name,
			Website:     company.Website,
			LastUpdated: utils.DateTime(company.UpdatedAt, constants.DateResponseFormat),
		},
		LastUpdated: utils.DateTime(project.UpdatedAt, constants.DateResponseFormat),
	}
}

func NewProjectController(services types.AllServices) ProjectController {
	return ProjectController{
		services: services,
		service:  services.ProjectService,
	}
}
