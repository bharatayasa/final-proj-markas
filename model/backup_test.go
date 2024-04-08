package model_test

import (
	"fmt"
	"testing"

	"github.com/bharatayasa/final-project/config"
	"github.com/bharatayasa/final-project/model"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("error, env not found")
	}

	config.OpenDb()
}

func TestGetLatest(t *testing.T) {
	Init()

	backupData := model.DatabaseBackup{
		File_name:     "mysql-2023-10-29-00-00-00-cv_kucing_oren-8634bf3f-23b5-45a7-8b78-fe9b1a3bcf66.sql.zip",
		Database_name: "db_5",
		File_path:     "haha/jsha.sajsa.zip",
	}

	_, err := backupData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := backupData.GetLatest(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	fmt.Println(res)
}

func TestGetByDbName(t *testing.T) {
	Init()

	backupData := model.DatabaseBackup{
		File_name:     "mysql-2023-10-29-00-00-00-cv_kucing_oren-8634bf3f-23b5-45a7-8b78-fe9b1a3bcf66.sql.zip",
		Database_name: "db_3",
		File_path:     "haha/jsha.sajsa.zip",
	}

	_, err := backupData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := backupData.GetByDbName(config.Mysql.DB, backupData.Database_name)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	fmt.Println(res)
}

func TestInsertData(t *testing.T) {
	Init()

	backupData := model.DatabaseBackup{
		File_name:     "mysql-2023-10-29-00-00-00-cv_kucing_oren-8634bf3f-23b5-45a7-8b78-fe9b1a3bcf66.sql.zip",
		Database_name: "db_3",
		File_path:     "haha/jsha.sajsa.zip",
	}

	_, err := backupData.InsertData(config.Mysql.DB)
	assert.Nil(t, err)
}
