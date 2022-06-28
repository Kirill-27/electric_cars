package handler

import (
	"github.com/Kirill-27/electric_cars/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		stations := api.Group("/stations")
		{
			stations.POST("/", h.createStation)
			stations.GET("/", h.getAllStations)
			stations.GET("/:id", h.getStationById)
			stations.PUT("/:id", h.updateStation)
			stations.DELETE("/:id", h.deleteStation)
		}
	}
	return router
}
