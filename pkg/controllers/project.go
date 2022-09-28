package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types"
	"gopher-camp/pkg/types/dto"
	"gopher-camp/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

type ProjectController struct {
	service  types.DOServiceProvider[models.Project]
	services types.AllServices
}

func (pc ProjectController) GetProjects(w http.ResponseWriter, _ *http.Request) {
	projects := pc.service.FindAll()
	var projectsResponse []dto.ProjectResponse
	for _, project := range projects {
		projectsResponse = append(projectsResponse, pc.convertToProject(project))
	}

	res, _ := json.Marshal(utils.SuccessArray(projectsResponse, "success", len(projects)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (pc ProjectController) DeleteProjectById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectId := params["projectId"]
	prjID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		pc.writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}

	deleted, err := pc.service.Delete(uint(prjID))
	if err != nil {
		pc.writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(utils.SingleObject(struct {
		Message string
		Deleted bool
	}{Message: "Deleted successfully", Deleted: deleted}))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (pc ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectId := params["projectId"]
	prjID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	project := models.NewProject()
	projectDTO := &dto.ProjectReqDTO{}
	utils.ParseBody(r, projectDTO)

	project = projectDTO.MapToDO(project)
	project, err = pc.service.Update(uint(prjID), project)

	if err != nil {
		pc.writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(utils.SingleObject(pc.convertToProject(*project)))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (pc ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	projectDO := models.NewProject()
	projectDTO := &dto.ProjectReqDTO{}
	utils.ParseBody(r, projectDTO)

	project, err := pc.service.Create(projectDTO.MapToDO(projectDO))
	if err != nil {
		pc.writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(utils.SingleObject(pc.convertToProject(*project)))
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (pc ProjectController) GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectId := params["projectId"]
	id, err := strconv.ParseUint(projectId, 0, 0)
	if err != nil {
		pc.writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}
	project, err := pc.service.FindById(uint(id))
	if err != nil {
		pc.writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(utils.SingleObject(pc.convertToProject(*project)))
	_, _ = w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func (pc ProjectController) writeFailureResponse(w http.ResponseWriter, err error, status int) {
	res, _ := json.Marshal(utils.CreateFailureWithMessage(err.(types.CustomError)))
	_, _ = w.Write(res)
	w.WriteHeader(status)
}

func (pc ProjectController) convertToProject(project models.Project) dto.ProjectResponse {
	company, _ := pc.services.CompanyService.FindById(project.CompanyID)
	company.Name = strings.Title(company.Name)

	return dto.ProjectResponse{
		ID:          project.ID,
		Deleted:     project.Deleted,
		Description: project.Description,
		Name:        strings.Title(project.Name),
		Company: dto.CompanyResponse{
			ID:          company.ID,
			Name:        company.Name,
			Website:     company.Website,
			Deleted:     company.Deleted,
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
