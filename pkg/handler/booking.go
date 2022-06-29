package handler

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllBookings(c *gin.Context) {
	stations, err := h.services.Booking.GetAllBookings()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stations)
}

func (h *Handler) getBookingById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	booking, err := h.services.Booking.GetBookingById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if booking == nil {
		newErrorResponse(c, http.StatusNotFound, "booking method with this id was not found")
		return
	}

	c.JSON(http.StatusOK, *booking)
}

func (h *Handler) deleteBooking(c *gin.Context) {
	isActiveAdmin, err := h.isActiveAdmin(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !isActiveAdmin {
		newErrorResponse(c, http.StatusUnauthorized, "you are not active admin")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	station, err := h.services.Booking.GetBookingById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if station == nil {
		newErrorResponse(c, http.StatusNotFound, "booking with this id to del was not found")
		return
	}

	err = h.services.Booking.DeleteBooking(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) createBooking(c *gin.Context) {
	customerId, err := getCustomerId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var booking data.Booking
	if err := c.BindJSON(&booking); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if customerId != booking.CustomerId {
		newErrorResponse(c, http.StatusBadRequest, "wrong customer id")
		return
	}

	id, err := h.services.Booking.CreateBooking(booking)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateBooking(c *gin.Context) {
	isActiveAdmin, err := h.isActiveAdmin(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !isActiveAdmin {
		newErrorResponse(c, http.StatusUnauthorized, "you are not active admin")
		return
	}
	
	c.Render(http.StatusNoContent, nil)
}
