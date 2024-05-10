package router

import (
	"api/handler/pinjaman"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/pinjaman", pinjaman.TambahPinjman)
	api.Get("/pinjaman", pinjaman.GetAllPinjaman)
	api.Get("/pinjaman/:nik", pinjaman.GetPinjamanByNIK)

}
