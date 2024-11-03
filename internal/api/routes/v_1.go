package routes

import (
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller"
)

func RouteV1(av *APIVersion) {
	userController := controller.NewUserController(av.contract.Service.User)
	authController := controller.NewAuthController(av.contract.Service.Auth)

	// auth routes
	auth := av.api.Group("/auth")

	auth.POST("/login", authController.Login)
	auth.POST("/new-access-token", authController.RefreshAccessToken)
	auth.POST("/logout", authController.Logout)

	// user routes
	user := av.api.Group("/users")
	// user.Use(middleware.JWTMiddleware()) // use middleware jwt general on user routes

	user.GET("", userController.Get)
	user.GET("/:id", userController.GetById)
	user.POST("", userController.Create)
	user.PUT("/:id", userController.Update)
	user.DELETE("/:id", userController.Delete)
}
