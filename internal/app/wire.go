package app

import (
	"uber-mini/internal/adapters/db"
	"uber-mini/internal/adapters/http"
	"uber-mini/internal/core/driver"

	"gorm.io/gorm"
)

func Init(dbConn *gorm.DB) *http.DriverHandler {
	driverRepository := db.NewDriverRepository(dbConn)
	driverService := driver.NewDriverService(driverRepository)
	driverHandler := http.DriverHandler{DriverService: driverService}
	return &driverHandler
}
