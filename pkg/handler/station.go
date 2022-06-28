package handler

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllStations(c *gin.Context) {

}

func (h *Handler) getStationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	station, err := h.services.Station.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, station)
}

func (h *Handler) createStation(c *gin.Context) {
	//TODO check is admin
	_, err := getCustomerId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var station data.Station
	if err := c.BindJSON(&station); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Station.Create(station)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteStation(c *gin.Context) {

}

func (h *Handler) updateStation(c *gin.Context) {

}
