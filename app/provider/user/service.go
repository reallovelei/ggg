package user

import "github.com/reallovelei/ggg/framework"

type UserService struct {
	container framework.Container
}

func NewUserService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &UserService{container: container}, nil
}

func (s *UserService) Foo() string {
	return ""
}
