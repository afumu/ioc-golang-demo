package main

import (
	"fmt"
	"github.com/alibaba/ioc-golang"
)

type UserService interface {
	GetUserName() string
}

type DaoService interface {
	GetUserNameById(id string) string
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type UserServiceImpl struct {
	DaoServiceImpl *DaoServiceImpl `singleton:""`
}

func (userServiceImpl UserServiceImpl) GetUserName() string {
	fmt.Println("获取用户名")
	return userServiceImpl.DaoServiceImpl.GetUserNameById("123")
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type DaoServiceImpl struct {
}

func (daoServiceImpl DaoServiceImpl) GetUserNameById(id string) string {
	fmt.Println("id:", id)
	return "zhangsan:" + id
}

func main() {
	// 加载所有结构
	if err := ioc.Load(); err != nil {
		panic(err)
	}

	userService, err := GetUserServiceImplSingleton()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userService.GetUserName())
}
