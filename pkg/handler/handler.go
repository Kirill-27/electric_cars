package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.singUp)
		auth.POST("sign-in", h.singIn)
	}

	api := router.Group("/api")
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
