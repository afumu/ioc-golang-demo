//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package controller

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &userController_{}
		},
	})
	userControllerStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &UserController{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(userControllerStructDescriptor)
}

type userController_ struct {
}

type UserControllerIOCInterface interface {
}

var _userControllerSDID string

func GetUserControllerSingleton() (*UserController, error) {
	if _userControllerSDID == "" {
		_userControllerSDID = util.GetSDIDByStructPtr(new(UserController))
	}
	i, err := singleton.GetImpl(_userControllerSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*UserController)
	return impl, nil
}

func GetUserControllerIOCInterfaceSingleton() (UserControllerIOCInterface, error) {
	if _userControllerSDID == "" {
		_userControllerSDID = util.GetSDIDByStructPtr(new(UserController))
	}
	i, err := singleton.GetImplWithProxy(_userControllerSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(UserControllerIOCInterface)
	return impl, nil
}