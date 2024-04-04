package model

type DatabaseBackup struct {
	Model
	File_name string `gorm:"not null" json:"file_name"`
	Timestamp string `gorm:"not null" json:"timestamp"`
}
