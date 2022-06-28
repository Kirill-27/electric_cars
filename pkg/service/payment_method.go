package service

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/Kirill-27/electric_cars/pkg/repository"
)

type PaymentMethodService struct {
	repo repository.PaymentMethod
}

func NewPaymentMethodService(repo repository.PaymentMethod) *PaymentMethodService {
	return &PaymentMethodService{repo: repo}
}

func (s *PaymentMethodService) Create(paymentMethod data.PaymentMethod) (int, error) {
	return s.repo.Create(paymentMethod)
}

func (s *PaymentMethodService) GetAll() ([]data.PaymentMethod, error) {
	return s.repo.GetAll()
}

func (s *PaymentMethodService) GetById(paymentMethodId int) (*data.PaymentMethod, error) {
	return s.repo.GetById(paymentMethodId)
}

func (s *PaymentMethodService) Delete(paymentMethodId int) error {
	return s.repo.Delete(paymentMethodId)
}

func (s *PaymentMethodService) Update(paymentMethod data.PaymentMethod) error {
	return s.repo.Update(paymentMethod)
}
