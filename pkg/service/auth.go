package service

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/Kirill-27/electric_cars/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateCustomer(customer data.Customer) (int, error) {
	return s.repo.CreateCustomer(customer)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {

	return "a", nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	return 1, nil
}

