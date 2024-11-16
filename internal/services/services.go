package services

import (
	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/models"
	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/storage"
)

type Services struct {
	SongService models.ZipperService
}

func ServiceInstance(storage *storage.Storage) *Services {
	return &Services{
		SongService: CreateZipperService(storage.ZipperStorage),
	}
}
