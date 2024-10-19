package routes

import "github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller"

func RouteV1(r *APIVersion) {
	userController := controller.NewUserController(r.contract.Service.User)

	user := r.api.Group("/users")
	user.GET("", userController.Get)
}
