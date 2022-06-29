package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllPaymentMethods(c *gin.Context) {
	paymentMethods, err := h.services.PaymentMethod.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, paymentMethods)
}

func (h *Handler) getPaymentMethodById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	paymentMethod, err := h.services.PaymentMethod.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if paymentMethod == nil {
		newErrorResponse(c, http.StatusNotFound, "payment method with this id was not found")
		return
	}

	c.JSON(http.StatusOK, *paymentMethod)
}

func (h *Handler) deletePaymentMethod(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	paymentMethod, err := h.services.PaymentMethod.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if paymentMethod == nil {
		newErrorResponse(c, http.StatusNotFound, "payment method with this id to del was not found")
		return
	}

	err = h.services.PaymentMethod.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
