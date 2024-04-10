package utils

import (
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
