package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage"
	"gopher-camp/pkg/types/dto"
	"gopher-camp/pkg/utils"
	"net/http"
	"strconv"
)

type ProjectController struct {
	service storage.Storage[models.Project]
}

func (pc ProjectController) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects := pc.service.FindAll()
	res, _ := json.Marshal(utils.SuccessArrayResponse(projects, "success"))
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
	projectDTO := &dto.ProjectDTO{}
	utils.ParseBody(r, projectDTO)
	project, _ := pc.service.Create(projectDTO)
	res, _ := json.Marshal(project)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}

func (pc ProjectController) GetProject(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//projectId := params["projectId"]
	//id, err := strconv.ParseInt(projectId, 0, 0)
	w.Header().Set("Content-Type", "application/json")
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
}

func NewProjectController(service storage.Storage[models.Project]) ProjectController {
	return ProjectController{service: service}
}
