package main

// list of functions to manipulate with User entity

type User struct{}

type UserService struct {
	dependencies any
}

func NewUserService(deps any) *UserService {
	return &UserService{dependencies: deps}
}

func (s *UserService) RegisterUser(data any) User {
	return User{}
}

func (s *UserService) BlockUser(id int64) {}

func (s *UserService) SetStatus(id int64, status string) {}

func (s *UserService) MarkAsDeleted(id int64) {}
