package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/bharatayasa/final-project/config"
	"github.com/bharatayasa/final-project/model"
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
