package http

import (
	"errors"
	"net/http"
	"strconv"
	"uber-mini/internal/core/driver"
	"uber-mini/internal/core/shared"

	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	DriverService *driver.DriverService
}

func (h *DriverHandler) handleGetDriver(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	drv, err := h.DriverService.Get(id)
	if err != nil {
		if errors.Is(err, shared.ErrDriverNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, drv)

}
