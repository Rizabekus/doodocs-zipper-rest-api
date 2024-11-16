package services

import "github.com/Rizabekus/doodocs-zipper-rest-api/internal/models"

type ZipperService struct {
	storage models.ZipperStorage
}

func CreateZipperService(storage models.ZipperService) *ZipperService {
	return &ZipperService{storage: storage}
}
