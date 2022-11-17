package controller

import "ioc-golang-demo/app/service"

// +ioc:autowire=true
// +ioc:autowire:type=singleton

// UserController 这是一个user
type UserController struct {
	UserService *service.UserService `singleton:""`
}

func (userController UserController) GetUserName() string {
	return userController.UserService.GetUserName()
}
