package http

import (
	"errors"
	"net/http"
	"strconv"
	"uber-mini/internal/core/driver"
	"uber-mini/internal/core/shared"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	driverService *driver.DriverService
	logger        *zap.Logger
}

func NewDriverHandler(driverService *driver.DriverService, logger *zap.Logger) *DriverHandler {
	return &DriverHandler{
		driverService: driverService,
		logger:        logger,
	}
}

func (h *DriverHandler) handleGetDriver(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		h.logger.Error("invalid id", zap.Int("id", id))
		return
	}
	drv, err := h.driverService.Get(id)
	if err != nil {
		if errors.Is(err, shared.ErrDriverNotFound) {
			h.logger.Warn("driver not found", zap.Int("id", id))
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.logger.Error("failed to get driver", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, drv)

}
