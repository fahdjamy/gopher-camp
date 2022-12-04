package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"profiler/pkg/constants"
	"profiler/pkg/models"
	"profiler/pkg/types"
	"profiler/pkg/types/dto"
	"profiler/pkg/utils"
	"profiler/pkg/utils/strings"
	"strconv"
)

type ProjectController struct {
	service  types.DOServiceProvider[models.Project]
	services types.AllServices
}

func (pc ProjectController) HandleGetAll(w http.ResponseWriter, req *http.Request) {
	projects := pc.service.FindAll()
	var projectsResponse []dto.ProjectResponse
	for _, project := range projects {
		projectsResponse = append(projectsResponse, pc.convertToProject(project))
	}

	res, _ := json.Marshal(utils.SuccessArray(projectsResponse, "success", len(projects)))
	writeJSONSuccessResponse(w, req, res, http.StatusOK)
}

func (pc ProjectController) HandleDeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectId := params["projectId"]
	prjID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}

	deleted, err := pc.service.Delete(uint(prjID))
	if err != nil {
		writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}

	res, _ := json.Marshal(utils.SingleObject(struct {
		Message string
		Deleted bool
	}{Message: "Deleted successfully", Deleted: deleted}))
	writeJSONSuccessResponse(w, r, res, http.StatusOK)
}

func (pc ProjectController) HandleUpdate(w http.ResponseWriter, r *http.Request) {
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
		writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(utils.SingleObject(pc.convertToProject(*project)))
	writeJSONSuccessResponse(w, r, res, http.StatusOK)
}

func (pc ProjectController) HandleCreate(w http.ResponseWriter, r *http.Request) {
	projectDO := models.NewProject()
	projectDTO := &dto.ProjectReqDTO{}
	utils.ParseBody(r, projectDTO)

	project, err := pc.service.Create(projectDTO.MapToDO(projectDO))
	if err != nil {
		writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(utils.SingleObject(pc.convertToProject(*project)))
	writeJSONSuccessResponse(w, r, res, http.StatusCreated)
}

func (pc ProjectController) HandleGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectId := params["projectId"]
	id, err := strconv.ParseUint(projectId, 0, 0)
	if err != nil {
		writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}
	project, err := pc.service.FindById(uint(id))
	if err != nil {
		writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(utils.SingleObject(pc.convertToProject(*project)))
	writeJSONSuccessResponse(w, r, res, http.StatusOK)
}

func (pc ProjectController) convertToProject(project models.Project) dto.ProjectResponse {
	company, _ := pc.services.CompanyService.FindById(project.CompanyID)
	company.Name = strings.Capitalize(company.Name)

	return dto.ProjectResponse{
		ID:          project.ID,
		Deleted:     project.Deleted,
		Description: project.Description,
		Name:        strings.Capitalize(project.Name),
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
