package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/dto"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage"
	"gopher-camp/pkg/utils"
	"net/http"
	"strconv"
)

type ProjectController struct {
	service storage.Storage[models.Project, dto.ProjectReqDTO]
}

func (pc ProjectController) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects := pc.service.FindAll()
	res, _ := json.Marshal(utils.SuccessArray(projects, "success"))
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
	//params := mux.Vars(r)
	//project := &models.Project{}
	//utils.ParseBody(r, project)
	//
	//projectId := params["projectId"]
	//_, err := strconv.ParseInt(projectId, 0, 0)
	//w.Header().Set("Content-Type", "application/json")
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
}

func (pc ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	projectDTO := &dto.ProjectReqDTO{}
	utils.ParseBody(r, projectDTO)
	project, err := pc.service.Create(projectDTO)
	if err != nil {
		res, _ := json.Marshal(utils.CreateFailure(err))
		_, _ = w.Write(res)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(utils.SingleObject(project))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}

func (pc ProjectController) GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectId := params["projectId"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		res, _ := json.Marshal(utils.CreateFailureWithMessage(err, "invalid id param"))
		_, _ = w.Write(res)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	project, err := pc.service.FindById(int(id))
	if err != nil {
		res, _ := json.Marshal(utils.CreateFailure(err))
		_, _ = w.Write(res)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	responseDTO := convertProjectToResponseDTO(*project)
	res, _ := json.Marshal(utils.SingleObject(responseDTO))
	_, _ = w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func convertProjectToResponseDTO(project models.Project) dto.ProjectResponseDTO {
	return dto.ProjectResponseDTO{
		Name:        project.Name,
		Description: project.Description,
		LastUpdated: utils.DateTime(project.UpdatedAt, constants.DateResponseFormat),
	}
}

func NewProjectController(service storage.Storage[models.Project, dto.ProjectReqDTO]) ProjectController {
	return ProjectController{service: service}
}
