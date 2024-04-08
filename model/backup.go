package model

import (
	"time"

	"gorm.io/gorm"
)

type DatabaseBackup struct {
	ID            uint      `gorm:"primarykey" json:"id"`
	Timestamp     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"timestamp"`
	File_name     string    `gorm:"not null" json:"file_name"`
	Database_name string    `gorm:"not null" json:"database_name"`
	File_path     string    `gorm:"not null" json:"file_path"`
}

func (bu *DatabaseBackup) Create(db *gorm.DB) (*DatabaseBackup, error) {
	err := db.
		Create(&bu).
		Error

	if err != nil {
		return nil, err
	}

	return bu, nil
}

func (bu *DatabaseBackup) InsertData(db *gorm.DB) (*DatabaseBackup, error) {
	err := db.
		Create(&bu).
		Error

	if err != nil {
		return nil, err
	}

	return bu, nil
}

func (bu *DatabaseBackup) GetLatest(db *gorm.DB) ([]DatabaseBackup, error) {
	var latestBackups []DatabaseBackup

	subquery := db.
		Select("MAX(timestamp) as timestamp, database_name").
		Group("database_name").
		Table("database_backups")

	err := db.
		Model(DatabaseBackup{}).
		Joins("JOIN (?) AS sub ON database_backups.timestamp = sub.timestamp AND database_backups.database_name = sub.database_name", subquery).
		Find(&latestBackups).
		Error

	if err != nil {
		return []DatabaseBackup{}, err
	}

	return latestBackups, nil
}

func (bu *DatabaseBackup) GetByDbName(db *gorm.DB, dbName string) ([]DatabaseBackup, error) {
	var backups []DatabaseBackup

	err := db.
		Where("database_name = ?", dbName).
		Find(&backups).
		Error
	if err != nil {
		return nil, err
	}

	return backups, nil
}
