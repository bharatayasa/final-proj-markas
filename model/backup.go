package model

type DatabaseBackup struct {
	Model
	File_name     string `gorm:"not null" json:"file_name"`
	Database_name string `gorm:"not null" json:"database_name"`
	File_path     string `gorm:"not null" json:"file_path"`
}
