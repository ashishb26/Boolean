package models

import "time"

// BoolTable is a struct that binds to the db table
type BoolTable struct {
	ID        string `json:"id" gorm:"primary key"`
	Value     bool   `json:"value"`
	Label     string `json:"label"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
