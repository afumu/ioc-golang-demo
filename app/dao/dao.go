package dao

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type UserDao struct {
}

func (userDao UserDao) GetUserName() string {
	return "zhangsan"
}
