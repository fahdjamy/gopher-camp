package routes

import (
	"profiler/pkg/controllers"
	"profiler/pkg/services"
	"profiler/pkg/types"
)

var ProfileRoutes = func(routers types.RestRouters, service services.ProfileService) {
	profileController := controllers.NewProfileController(service)

	routers.SetSubRoutePrefix("/profile")
	routers.Get("/profile", profileController.GetOwnProfile)
}
