package http

import (
	"strconv"
	"uber-mini/internal/core/driver"

	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	DriverService *driver.DriverService
}

func (h *DriverHandler) handleGetDriver(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"}) //i need a validator for every parameter like java
	}
	driver, err := h.DriverService.Get(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()}) //manage errors better
		return
	}
	ctx.JSON(200, driver)

}
