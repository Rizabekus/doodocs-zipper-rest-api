package storage

import (
	"database/sql"

	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/models"
)

type Storage struct {
	ZipperStorage models.ZipperStorage
}

func StorageInstance(db *sql.DB) *Storage {
	return &Storage{ZipperStorage: CreateZipperStorage(db)}
}
