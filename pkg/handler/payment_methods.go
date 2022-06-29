package handler

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createPaymentMethod(c *gin.Context) {
	customerId, err := getCustomerId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var paymentMethod data.PaymentMethod
	if err := c.BindJSON(&paymentMethod); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if customerId != paymentMethod.CustomerId {
		newErrorResponse(c, http.StatusBadRequest, "wrong customer id")
		return
	}

	id, err := h.services.PaymentMethod.Create(paymentMethod)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

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
	customerId, err := getCustomerId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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
	if paymentMethod.CustomerId != customerId {
		newErrorResponse(c, http.StatusBadRequest, "you can delete this")
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

func (h *Handler) updatePaymentMethod(c *gin.Context) {
	customerId, err := getCustomerId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	paymentMethog, err := h.services.PaymentMethod.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if paymentMethog == nil {
		newErrorResponse(c, http.StatusNotFound, "station with this id to upd was not found")
		return
	}

	if err := c.BindJSON(&paymentMethog); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if customerId != paymentMethog.CustomerId {
		newErrorResponse(c, http.StatusBadRequest, "wrong customer id")
		return
	}

	paymentMethog.Id = id

	err = h.services.PaymentMethod.Update(*paymentMethog)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Render(http.StatusNoContent, nil)
}
