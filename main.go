package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/eko-gsi/keuangan-be/internal/database"
	"github.com/eko-gsi/keuangan-be/internal/handler"
	"github.com/eko-gsi/keuangan-be/internal/model"
)

func main() {
	_ = godotenv.Load()

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "keuangan.db"
	}

	db, err := database.Connect(dbPath)
	if err != nil {
		log.Fatal("gagal koneksi db: ", err)
	}
	if err := model.Migrate(db); err != nil {
		log.Fatal("gagal migrasi: ", err)
	}

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	th := &handler.TransaksiHandler{DB: db}
	app.Post("/api/v1/transaksi", th.Create)

	log.Fatal(app.Listen(":3000"))
}
