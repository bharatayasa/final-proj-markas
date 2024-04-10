package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bharatayasa/final-project/model"
	"github.com/bharatayasa/final-project/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetDatabaseLatest(c *fiber.Ctx) error {
	latestData, err := utils.GetLatestUtils()
	if err != nil {
		logrus.Error("Error on get latest utils: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server Error",
		})
	}

	latestBackups := make([]map[string]interface{}, 0)

	for _, data := range latestData {
		latestBackup := map[string]interface{}{
			"database_name": data.Database_name,
			"latest_backup": map[string]interface{}{
				"id":        data.ID,
				"file_name": data.File_name,
				"timestamp": data.Timestamp,
			},
		}
		latestBackups = append(latestBackups, latestBackup)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    latestBackups,
		"message": "Success",
	})
}

func GetDatabaseByDbName(c *fiber.Ctx) error {
	dbName := c.Params("database_name")

	if dbName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Database name is required",
		})
	}

	backups, err := utils.GetByDbNametUtils(dbName)
	if err != nil {
		logrus.Error("Error on get backups by database name: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server Error",
		})
	}

	histories := make([]map[string]interface{}, 0)

	for _, data := range backups {
		history := map[string]interface{}{
			"id":        data.ID,
			"file_name": data.File_name,
			"timestamp": data.Timestamp,
		}
		histories = append(histories, history)
	}

	response := fiber.Map{
		"data": fiber.Map{
			"database_name": dbName,
			"histories":     histories,
		},
		"message": "success",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func InsertNewData(c *fiber.Ctx) error {
	dbName := c.Params("database_name")

	if dbName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Database name is required",
		})
	}

	file, err := c.FormFile("file_name")
	if err != nil {
		logrus.Error("Error getting file:", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get file",
		})
	}

	filePath := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		logrus.Error("Error saving file:", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save file",
		})
	}

	backupData := model.DatabaseBackup{
		Timestamp:     time.Now(),
		File_name:     file.Filename,
		Database_name: dbName,
		File_path:     filePath,
	}

	insertedData, err := utils.InsertDataUtils(&backupData)
	if err != nil {
		logrus.Error("Error inserting data using utils:", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to insert data using utils",
		})
	}

	response := fiber.Map{
		"data": fiber.Map{
			"id":            insertedData.ID,
			"database_name": insertedData.Database_name,
			"file_name":     insertedData.File_name,
			"timestamp":     insertedData.Timestamp,
		},
		"message": "success",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func DownloadFile(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	type DownloadRequest struct {
		FilePath string `json:"file_path"`
	}
	var requestBody DownloadRequest
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
		})
	}

	if requestBody.FilePath == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "File path is required in the request body",
		})
	}

	backupData, err := utils.DownloadFileUtils(uint(id), requestBody.FilePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to download file",
		})
	}

	if backupData == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "File not found for the given ID and file path",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File successfully retrieved",
	})
}
