package service

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/Kirill-27/electric_cars/pkg/repository"
)

type BookingService struct {
	repo repository.Booking
}

func NewBookingService(repo repository.Booking) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) CreateBooking(booking data.Booking) (int, error) {
	return s.repo.CreateBooking(booking)
}

func (s *BookingService) GetAllBookings() ([]data.Booking, error) {
	return s.repo.GetAllBookings()
}

func (s *BookingService) GetBookingById(bookingId int) (*data.Booking, error) {
	return s.repo.GetBookingById(bookingId)
}

func (s *BookingService) DeleteBooking(bookingId int) error {
	return s.repo.DeleteBooking(bookingId)
}

func (s *BookingService) UpdateBooking(booking data.Booking) error {
	return s.repo.UpdateBooking(booking)
}
