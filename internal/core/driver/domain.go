package driver

import (
	"uber-mini/internal/core/ride"
)

type Driver struct { //Domain depends of GORM
	ID     uint `gorm:"primary_key"`
	Name   string
	Email  string `gorm:"unique"`
	Gender string
	CarID  uint
	Car    Car
	Ride   []ride.Ride `gorm:"foreignKey: DriverID"`
	Rating float32
}
type Car struct { //Domain depends of GORM
	ID    uint   `gorm:"primary_key"`
	Plate string `gorm:"unique"`
	Model string
	Color string
}
