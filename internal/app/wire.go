package app

import (
	"uber-mini/internal/adapters/db"
	"uber-mini/internal/adapters/http"
	"uber-mini/internal/core/driver"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

func Init(dbConn *gorm.DB) *http.DriverHandler {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	driverRepository := db.NewDriverRepository(dbConn)
	driverService := driver.NewDriverService(driverRepository, logger)
	driverHandler := http.NewDriverHandler(driverService, logger)

	return driverHandler
}
