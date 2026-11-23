package model

import "database/sql"

type Transaksi struct {
	ID         int64   `json:"id"`
	Tipe       string  `json:"tipe"` // "masuk" atau "keluar"
	Keterangan string  `json:"keterangan"`
	Jumlah     float64 `json:"jumlah"`
	Tanggal    string  `json:"tanggal"` // format YYYY-MM-DD
}

func Migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS transaksis (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tipe TEXT NOT NULL CHECK(tipe IN ('masuk','keluar')),
		keterangan TEXT NOT NULL,
		jumlah REAL NOT NULL CHECK(jumlah > 0),
		tanggal TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}
