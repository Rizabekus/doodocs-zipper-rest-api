package services

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/models"
	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/custom_errors"
)

type ZipperService struct{}

func CreateZipperService() *ZipperService {
	return &ZipperService{}
}
func (zs *ZipperService) GetArchiveInfo(fileBytes []byte, fileSize int64) (models.ArchiveInfo, error) {
	archive, err := zip.NewReader(bytes.NewReader(fileBytes), fileSize)
	if err != nil {
		return models.ArchiveInfo{}, fmt.Errorf("failed to read archive: %w", err)
	}

	var files []models.File
	totalFiles := 0
	totalSize := float64(0)

	for _, f := range archive.File {
		totalFiles++
		totalSize += float64(f.UncompressedSize64)

		rc, err := f.Open()
		if err != nil {
			return models.ArchiveInfo{}, fmt.Errorf("failed to open file %s in archive: %w", f.Name, err)
		}

		buffer := make([]byte, 512)
		n, _ := rc.Read(buffer)
		mimeType := http.DetectContentType(buffer[:n])
		rc.Close()

		files = append(files, models.File{
			File_Path: f.Name,
			Size:      float64(f.UncompressedSize64),
			MIMEType:  mimeType,
		})
	}

	return models.ArchiveInfo{
		FileName:     "",
		Archive_Size: float64(fileSize),
		Total_Size:   totalSize,
		Total_Files:  float64(totalFiles),
		Files:        files,
	}, nil
}
func (zs *ZipperService) CreateArchive(files []*multipart.FileHeader) (bytes.Buffer, error) {
	validMimeTypes := map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/xml": true,
		"image/jpeg":      true,
		"image/png":       true,
	}

	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return buf, err
		}
		defer file.Close()

		fileBuffer := make([]byte, 512)
		_, err = file.Read(fileBuffer)
		if err != nil {
			return buf, err
		}
		mimeType := http.DetectContentType(fileBuffer)

		if !validMimeTypes[mimeType] {

			return buf, custom_errors.ErrWrongMIMEType
		}

		zipFileWriter, err := zipWriter.Create(fileHeader.Filename)
		if err != nil {
			return buf, err
		}

		_, err = io.Copy(zipFileWriter, file)
		if err != nil {
			return buf, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return buf, err
	}
	return buf, nil
}
