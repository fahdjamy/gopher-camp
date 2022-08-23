package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/utils"
	"net/http"
	"strconv"
)

var dbProject models.Project

func GetProjects(w http.ResponseWriter, r *http.Request) {
	projects := dbProject.GetAll()
	res, _ := json.Marshal(projects)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}

func DeleteProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectId := params["projectId"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	project := dbProject.DeleteById(id)
	if project.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(struct {
		message string
	}{
		message: "Project deleted",
	})
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	project := &models.Project{}
	utils.ParseBody(r, project)

	projectId := params["projectId"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	projectDetails, db := dbProject.FindById(id)
	if projectDetails.ID == "" && projectDetails.ID != projectId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	projectDetails.Name = project.Name
	projectDetails.Description = project.Description
	db.Save(&projectDetails)

	res, _ := json.Marshal(projectDetails)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	project := &models.Project{}
	utils.ParseBody(r, project)
	savedProject := project.Create()
	res, _ := json.Marshal(savedProject)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectId := params["projectId"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbProject, _ := dbProject.FindById(id)
	res, _ := json.Marshal(dbProject)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}
