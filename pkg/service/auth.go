package service

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/Kirill-27/electric_cars/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "dfvjdfv@21#@d(q*Djdsdf"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	CustomerId int `json:"customer_id"`
}

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
	customer, err := s.repo.GetCustomer(username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		customer.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	return 1, nil
}
