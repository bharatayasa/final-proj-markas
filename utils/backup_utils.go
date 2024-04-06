package utils

import (
	"github.com/bharatayasa/final-project/config"
	"github.com/bharatayasa/final-project/model"
)

func GetLatestUtils() ([]model.DatabaseBackup, error) {
	var Backup model.DatabaseBackup

	return Backup.GetLatest(config.Mysql.DB)
}