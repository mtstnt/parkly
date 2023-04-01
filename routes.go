package main

import (
	"parkly/handlers"
	"parkly/services"

	"github.com/gin-gonic/gin"
)

var (
	userHandler *handlers.User
)

func setupDependencies() error {
	userService := services.NewUser()
	userHandler = handlers.NewUser(userService)

	return nil
}

func setupRoutes(router gin.IRouter) error {
	router.GET("/", userHandler.GetAllUsers)
	return nil
}
