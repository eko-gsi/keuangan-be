package handler

import (
	"database/sql"

	"github.com/eko-gsi/keuangan-be/internal/model"
	"github.com/gofiber/fiber/v2"
)

type TransaksiHandler struct {
	DB *sql.DB
}

func (h *TransaksiHandler) Create(c *fiber.Ctx) error {
	var input model.Transaksi
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "body tidak valid"})
	}

	if input.Tipe != "masuk" && input.Tipe != "keluar" {
		return c.Status(400).JSON(fiber.Map{"error": "tipe harus 'masuk' atau 'keluar'"})
	}
	if input.Keterangan == "" || input.Jumlah <= 0 || input.Tanggal == "" {
		return c.Status(400).JSON(fiber.Map{"error": "keterangan, jumlah (>0), dan tanggal wajib diisi"})
	}

	res, err := h.DB.Exec(
		`INSERT INTO transaksis (tipe, keterangan, jumlah, tanggal) VALUES (?, ?, ?, ?)`,
		input.Tipe, input.Keterangan, input.Jumlah, input.Tanggal,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	id, _ := res.LastInsertId()
	input.ID = id
	return c.Status(201).JSON(input)
}
