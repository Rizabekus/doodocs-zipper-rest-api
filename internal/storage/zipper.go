package storage

import "database/sql"

type ZipperDB struct {
	DB *sql.DB
}

func CreateZipperStorage(db *sql.DB) *ZipperDB {
	return &ZipperDB{DB: db}
}
