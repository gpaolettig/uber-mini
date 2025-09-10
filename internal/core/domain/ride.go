package domain

import "time"

type Ride struct {
	ID          uint `gorm:"primary_key"`
	Origin      string
	Destination string
	FareCents   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
