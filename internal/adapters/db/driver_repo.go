package db

import (
	"uber-mini/internal/core/driver"

	"gorm.io/gorm"
)

type driverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) *driverRepository {
	return &driverRepository{db: db}
}
func (r *driverRepository) FindById(id int) (driver.Driver, error) {
	var drv driver.Driver
	if err := r.db.First(&drv, id).Error; err != nil {
		return driver.Driver{}, err
	}
	return drv, nil
}
