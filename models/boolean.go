package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

// BoolRecord is a struct that is used to create and
// bind to the database table that stores the boolean
type BoolRecord struct {
	ID        string    `json:"id" gorm:"primary key"`
	Value     bool      `json:"value"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AddNewRecord adds a new record to the database
func (record *BoolRecord) AddNewRecord(DB *gorm.DB) {
	record.ID = xid.New().String()
	DB.Create(record)
}

// GetRecordByID extracts and returns a boolean record given the ID
func (record *BoolRecord) GetRecordByID(DB *gorm.DB, recordID string) error {
	if err := DB.Where("id=?", recordID).First(record).Error; err != nil {
		return err
	}

	return nil
}

// UpdateRecord updates a database record
func (record *BoolRecord) UpdateRecord(DB *gorm.DB) error {
	DB.Model(&record).Updates(record)
	return nil
}

// DeleteRecordByID deletes a boolean record from the database given the id of the
// database record
func DeleteRecordByID(DB *gorm.DB, recordID string) error {

	var record BoolRecord

	if err := DB.Where("id=?", recordID).First(&record).Error; err != nil {
		return err
	}

	DB.Delete(&record)
	return nil
}
