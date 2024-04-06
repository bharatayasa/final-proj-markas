package controller

import (
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

// func AddBackupData(c *fiber.Ctx) error {
// 	var data model.DatabaseBackup

// 	if err := c.BodyParser(&data); err != nil {
// 		logrus.Error("Error on parsing cars data: ", err.Error())
// 		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "Server Error",
// 		})
// 	}

// 	err := utils.InsertBackupData(data)

// 	if err != nil {
// 		logrus.Error("Error on insert cars data: ", err.Error())
// 		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "Server Error",
// 		})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 		"data":    data,
// 		"message": "Success",
// 	})
// }
