package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/bharatayasa/final-project/config"
	"github.com/bharatayasa/final-project/model"
	"gorm.io/gorm"
)

func GetLatestUtils() ([]model.DatabaseBackup, error) {
	var Backup model.DatabaseBackup

	return Backup.GetLatest(config.Mysql.DB)
}

func GetByDbNametUtils(dbName string) ([]model.DatabaseBackup, error) {
	var Backup model.DatabaseBackup

	return Backup.GetByDbName(config.Mysql.DB, dbName)
}

func InsertDataUtils(backup *model.DatabaseBackup) (*model.DatabaseBackup, error) {

	return backup.InsertData(config.Mysql.DB)
}

func DownloadFileUtils(ID uint, File_path string) (*model.DatabaseBackup, error) {
	var backup model.DatabaseBackup

	backupData, err := (&backup).DownloadFile(config.Mysql.DB, ID, File_path)
	if err != nil {
		return nil, err
	}

	return backupData, nil
}

func GetFilePathByID(id uint) (string, error) {
	var backup model.DatabaseBackup

	backupData, err := backup.GetById(config.Mysql.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("no file path found for the given ID: %d", id)
		}
		return "", fmt.Errorf("failed to retrieve file path for the given ID: %w", err)
	}

	if len(backupData) == 0 {
		return "", fmt.Errorf("no file path found for the given ID: %d", id)
	}

	return backupData[0].File_path, nil
}

func MoveFileUtils(sourcePath, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(filepath.Join(destinationPath, filepath.Base(sourcePath)))
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	return nil
}
