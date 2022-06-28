package service

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/Kirill-27/electric_cars/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user data.Customer) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		//Authorization: new(Authorization),
		TodoList: new(TodoList),
		TodoItem: new(TodoItem),
	}
}