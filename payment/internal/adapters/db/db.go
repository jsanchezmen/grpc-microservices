package db

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderId    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}
