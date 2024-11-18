package models

import (
	"bytes"
	"mime/multipart"
)

type ZipperService interface {
	GetArchiveInfo(fileBytes []byte, fileSize int64) (ArchiveInfo, error)
	CreateArchive(files []*multipart.FileHeader) (bytes.Buffer, error)
}

type ArchiveInfo struct {
	FileName     string  `json:"filename"`
	Archive_Size float64 `json:"archive_size"`
	Total_Size   float64 `json:"total_size"`
	Total_Files  float64 `json:"total_files"`
	Files        []File  `json:"files"`
}

type File struct {
	File_Path string  `json:"file_path"`
	Size      float64 `json:"size"`
	MIMEType  string  `json:"mimetype"`
}
type Response struct {
	Message string `json:"message"`
}
