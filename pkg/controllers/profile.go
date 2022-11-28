package controllers

import (
	"encoding/json"
	"net/http"
	"profiler/pkg/services"
	"profiler/pkg/utils"
)

type ProfileController struct {
	service services.ProfileService
}

func (pc ProfileController) GetOwnProfile(w http.ResponseWriter, r *http.Request) {
	profile, err := pc.service.FindById(1)
	if err != nil {
		writeFailureResponse(w, err, http.StatusBadRequest)
		return
	}

	res, _ := json.Marshal(utils.SingleObject(*profile))
	writeJSONSuccessResponse(w, r, res, http.StatusOK)
}

func NewProfileController(service services.ProfileService) ProfileController {
	return ProfileController{
		service: service,
	}
}
