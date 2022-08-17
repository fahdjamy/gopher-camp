package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Company     *Company `json:"company"`
}

type Company struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Founder string `json:"founder"`
	Website string `json:"website"`
}

var projects []Project

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	projects = append(projects, Project{
		ID:          "1",
		Name:        "Fyipe",
		Description: "Random",
		Company: &Company{
			ID:      "1",
			Founder: "Nawaz",
			Website: "https://google.com",
		},
	})
	err := json.NewEncoder(w).Encode(projects)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func DeleteProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, prjt := range projects {
		if prjt.ID == params["id"] {
			projects = append(projects[:index], projects[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(projects)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, prjt := range projects {
		if prjt.ID == params["id"] {
			projects = append(projects[:index], projects[index+1:]...)
		}
	}
	var project Project
	project.ID = params["id"]
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		log.Fatal(err)
		return
	}
	projects = append(projects, project)

	err = json.NewEncoder(w).Encode(project)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var project Project

	_ = json.NewDecoder(r.Body).Decode(&project)
	project.ID = strconv.Itoa(1000000000000)
	projects = append(projects, project)

	err := json.NewEncoder(w).Encode(project)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetOneProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, prjt := range projects {
		if prjt.ID == params["id"] {
			err := json.NewEncoder(w).Encode(prjt)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	//err := json.NewEncoder(w).Encode(projects)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
}
