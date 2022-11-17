package main

import (
	"fmt"
	"github.com/alibaba/ioc-golang"
	"ioc-golang-demo/app/controller"
)

func main() {
	// 加载所有结构
	if err := ioc.Load(); err != nil {
		panic(err)
	}

	userController, err := controller.GetUserControllerSingleton()
	if err != nil {
		fmt.Println(err)
	}
	name := userController.GetUserName()
	fmt.Println(name)
}
