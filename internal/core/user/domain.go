package user

import (
	"time"
	"uber-mini/internal/core/ride"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Email     string `gorm:"unique"`
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Rides     []ride.Ride `gorm:"foreignKey: UserID"`
}
