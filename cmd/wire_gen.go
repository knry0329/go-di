// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/knry0329/go-di/controller"
	"github.com/knry0329/go-di/repository"
	"github.com/knry0329/go-di/service"
)

// Injectors from wire.go:

func initializeUser() controller.UserController {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController
}
