package handler

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllStations(c *gin.Context) {
	stations, err := h.services.Station.GetAllStations()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stations)
}

func (h *Handler) getStationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	station, err := h.services.Station.GetStationById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if station == nil {
		newErrorResponse(c, http.StatusNotFound, "station with this id was not found")
		return
	}

	c.JSON(http.StatusOK, *station)
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

	id, err := h.services.Station.CreateStation(station)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteStation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	station, err := h.services.Station.GetStationById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if station == nil {
		newErrorResponse(c, http.StatusNotFound, "station with this id to del was not found")
		return
	}

	err = h.services.Station.DeleteStation(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) updateStation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	station, err := h.services.Station.GetStationById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if station == nil {
		newErrorResponse(c, http.StatusNotFound, "station with this id to upd was not found")
		return
	}

	if err := c.BindJSON(&station); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	station.Id = id

	err = h.services.Station.UpdateStation(*station)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Render(http.StatusNoContent, nil)
}
