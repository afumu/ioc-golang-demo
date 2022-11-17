package service

import "ioc-golang-demo/app/dao"

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type UserService struct {
	UserDao *dao.UserDao `singleton:""`
}

func (userService UserService) GetUserName() string {
	return userService.UserDao.GetUserName()
}
