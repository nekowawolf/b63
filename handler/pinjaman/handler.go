package pinjaman

import (
	"api/model"
	"api/repository/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func TambahPinjman(c *fiber.Ctx) error {
	var requestData model.Pinjamann
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	requestData.ID = uuid.New().String()

	// If data doesn't exist, proceed with insertion
	if err := db.InsertPinjaman(requestData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data into the database",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data inserted successfully",
		"data":    requestData,
	})
}

func GetAllPinjaman(c *fiber.Ctx) error {
	filter := bson.M{}

	requestData, err := db.GetAllPinjaman(filter)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    requestData,
	})
}

func GetPinjamanByNIK(c *fiber.Ctx) error {
	nik := c.Params("nik")
	filter := bson.M{"nik": nik}
	requestData, err := db.GetOnePinjaman(filter)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    requestData,
	})
}
