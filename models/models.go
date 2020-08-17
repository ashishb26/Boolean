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

// InputBool is used to bind to the input JSON file as the InputBool
// only has value and label
type InputBool struct {
	Value bool   `json:"value"`
	Label string `json:"label"`
}
