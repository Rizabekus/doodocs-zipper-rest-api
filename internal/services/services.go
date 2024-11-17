package services

import (
	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/models"
)

type Services struct {
	ZipperService models.ZipperService
}

func ServiceInstance() *Services {
	return &Services{
		ZipperService: CreateZipperService(),
	}
}
