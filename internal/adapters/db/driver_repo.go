package db

import (
	"uber-mini/internal/core/driver"
	"uber-mini/internal/core/shared"

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
		if err == gorm.ErrRecordNotFound {
			return driver.Driver{}, shared.ErrDriverNotFound
		}
		return driver.Driver{}, err
	}
	return drv, nil
}
