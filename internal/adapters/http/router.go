package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, driverHandler *DriverHandler) {
	r.GET(("/drivers/:id"), driverHandler.handleGetDriver)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "UP",
		})
	})
}
